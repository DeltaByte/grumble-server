package middleware

import (
	"time"

	"github.com/grumblechat/server/internal/config"
	"github.com/grumblechat/server/internal/logging"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Logger(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// capture initial req/res
			req := ctx.Request()
			res := ctx.Response()
			start := time.Now()

			// wait for request to complete
			err := next(ctx)
			if err != nil {
				ctx.Error(err)
			}

			// capture end time and setup logger
			stop := time.Now()
			log := logging.Request()
			defer log.Sync()

			// request ID
			requestID := req.Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = res.Header().Get(echo.HeaderXRequestID)
			}

			// bytes in
			bytesIn := req.Header.Get(echo.HeaderContentLength)
			if bytesIn == "" {
				bytesIn = "0"
			}

			// main fields (avoid anything sensitive here)
			fields := []zap.Field{
				zap.Time("time", time.Now()),
				zap.String("host", req.Host),
				zap.String("requestId", requestID),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.Int("status", res.Status),
				zap.Error(err),
				zap.Duration("latency", stop.Sub(start)),
				zap.String("latency_human", stop.Sub(start).String()),
				zap.String("bytes_in", bytesIn),
				zap.Int64("bytes_out", res.Size),
			}

			// only log 'remote_ip' if enabled
			if cfg.Logging.RemoteIP {
				fields = append(fields, zap.String("remote_ip", ctx.RealIP()))
			}

			// only log 'user_agent' if enabled
			if cfg.Logging.UserAgent {
				fields = append(fields, zap.String("user_agent", req.UserAgent()))
			}

			// log request
			log.Info("Request", fields...)
			return nil
		}
	}
}
