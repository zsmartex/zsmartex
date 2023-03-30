package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/looplab/eventhorizon/uuid"
	"github.com/zsmartex/pkg/v3/infrastucture/redis"
	"github.com/zsmartex/zsmartex/cmd/proxy/config"
	"github.com/zsmartex/zsmartex/pkg/session"
	userv1 "github.com/zsmartex/zsmartex/proto/api/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func newGateway(ctx context.Context, cfg *config.Config, opts ...runtime.ServeMuxOption) (*runtime.ServeMux, error) {
	userEndpoint := fmt.Sprintf("%s:%d", cfg.GRPC.UserHost, cfg.GRPC.UserPort)

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func allowCORS(c *gin.Context) {
	if origin := c.Request.Header.Get("Origin"); origin != "" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		if c.Request.Method == "OPTIONS" && c.Request.Header.Get("Access-Control-Request-Method") != "" {
			preflightHandler(c.Writer, c.Request)

			return
		}
	}
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	headers := []string{"*"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("Config error: %s", err)
	}

	server := gin.New()
	server.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	server.Use(allowCORS)

	redisClient, err := redis.New(cfg.Redis.URI)
	if err != nil {
		logger.Error("failed connect to redis", err)
	}

	sessionStore := session.NewStore(redisClient)

	gw, err := newGateway(
		ctx,
		cfg,
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := metadata.Pairs("user-agent", r.Header.Get("User-Agent"))

			sessionCookie, err := r.Cookie(sessionStore.CookieKey)
			if err == nil {
				md.Append("session_id", sessionCookie.Value)
			}

			return md
		}),
		runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, m proto.Message) error {
			md, ok := runtime.ServerMetadataFromContext(ctx)
			if !ok {
				return nil
			}

			if len(md.HeaderMD.Get("session_id")) == 0 {
				return nil
			}

			sessionID, err := uuid.Parse(md.HeaderMD.Get("session_id")[0])
			if err != nil {
				return nil
			}

			session, err := sessionStore.GetSession(ctx, sessionID)
			if err != nil {
				return err
			}

			session.SetSession(ctx, w)

			return nil
		}),
	)
	if err != nil {
		logger.Error("failed to create a new gateway", err)
	}

	server.Group("*{grpc_gateway}").Any("", func(c *gin.Context) {
		gw.ServeHTTP(c.Writer, c.Request)
	})

	go func() {
		<-ctx.Done()
		logger.Info("shutting down the http server")
	}()

	logger.Info("start listening...", "address", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err := server.Run(fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)); err != nil {
		logger.Error("failed to listen and serve", err)
	}
}
