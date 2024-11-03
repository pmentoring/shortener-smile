package main

import (
	"log/slog"
	"os"
	"os/signal"
	"shortener-smile/database"
	"shortener-smile/internal/app/grpc"
	"shortener-smile/internal/common"
	"shortener-smile/internal/shortener/repository"
	"shortener-smile/internal/shortener/service"
	"syscall"
)

func main() {
	log := createLogger()

	log.Info("Booting app")

	db, err := database.GetConnection()

	if err != nil {
		log.Error(err.Error())
	}

	appCtx := getAppContext()

	shortenerService := service.NewShortenLinkService(repository.NewLinksRepository(db), appCtx, log)
	unshortenerService := service.NewUnshortenLinkService(repository.NewLinksRepository(db), log)

	app := grpc.NewApp(log, 8000, shortenerService, unshortenerService)

	go app.MustStart()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Info("Shutting down...")
	app.GracefulShutdown()
}

func getAppContext() *common.ApplicationContext {
	//return common.NewApplicationContext(os.Getenv("INSTANCE_ID"), os.Getenv("APP_BASE_URL"))
	return common.NewApplicationContext("01", "http://localhost:8000/", "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAoj+0uMh2G3Dg2cl+U3Ph1tsIdw9VpiP9grdprOQiyTwbrEbk\n4P6w1WC2nHzTIItLdcsfhirBtpmE29hOS6OrD52mgsvm06R2fmYLU23DwSw2bjwN\n7HY141nWDlfo1ZxOy5aWXshZYhSyYymBe963ib3u1PyIrQgCG1kpIYSe4wOkpaun\nGrHxpFznw0T7HLZtZ0H9n4P9Rb73QSCBmBjdhNKhM7hzRAnYW5eqkMavAOIM415m\nNeshMPzlPNY5ZBah2oi2Dec4cD4NHo6IEA8iaZFSdPTHsszZ0rebvw0LC/ugzmSe\nrLEzeCCQnbqmcHm0Yz46MmWQcS9/rUbDsw9eywIDAQABAoIBAFivFatC9zu8kUq5\nljErUDBOfOwHF7u+RF1kv4BfGpyAHGgoQAazcBjRyj7zA7vwJ4j1Q//j8wpKE5LP\nQsyZY3EfrNoldcANTy97u4kYoD1AW2gZr0qmuoHQou4oIv5Pd8pUP1EDlcfEpFUL\nZNk+PXgutxL4DgT2MaCdNZw6j1Zp0B89jHdAZg8ueSIs9xuapBdFcRu+dNMPAnUL\n/oAEdYkE+rl5j+h8QAEoAgkJ2wzJ7gO3ewkI6i7jPC2KsKzxndCZHg5qb/PP15ZW\nFpeshahmneSzhsdpciBbT55vHWiR1OiiYtV77HHATx08c6vyzC6yL6kKA5+GqP1/\nwWx6V0ECgYEA7F3wSr8bUgPJymDh021bBNE/iqwiGHD5lNulSGGekGXh23MBN3nW\nUoDRtT3H89tyMU1uBYMLjyFR0DKDawIg+1NMNeS7bOLD6/OGqFNrnAZJXajku0OG\n8Z0a1tt667/rVnFrwz7eVgOQrtftAyz24dcWnUUm+OQlNLOkp7Hn3vsCgYEAr7m7\n/ePsDp1TjVJV8HBbFPzLCGODZHEJUmqjSioB6liuHLxwrIgnFaU0E5bNIlJwMe58\nr4jd6wzXqxwymk1N46+fEvPBSCgDZx3BRlU8a7mYCYbVAiMXITX87o33JUJvhxrJ\naJYnH0bDIUslD5YxKo1sloPHkgIycBtfmRR3NnECgYEAg/3NMqBXbEwrQjUpxjw5\n6I3E8vY+r/Feng6t1KaWukH6huZI1qbV1QqUkKYoY3e24+s+NhvAt7O+kQm0M2xj\n6JTSaQIK27oh+D/sgVDcNauZCpFG8X+6+m9lznePw9hPXd1GdShFVjj1cX2on+5b\nCzBmM3qrUa9tS2nJ189UEUMCgYBqtrlyYpJ9AY+219R2slSttyK5Um19dqO5mCbX\nqBEHtpjtvIYUUI6jt3Kbns2uVS/rMrAnjU9vKZpjQDl+uQlfBszsnni3EA2Y9lnk\nE2kg5lFtH6OFq7elJHsYh8AZsXG3M04ypFTXyoo/UroL8CKle8MA2MA+a/UjXYw9\nk08eEQKBgCQXbUmRYltuixWJ6AfMtfAUvqjuG4Jl2Pa8X07zNCBJL+Y2ZrCUnNUD\n5zd8XAFi5O0W+LEBZAjWyMsbfEyzAOY9fQiuEViHS8cQr8Gu/fjFaxKYl+zs/pOR\n32CNHM4Pb5DQonzkysYLS7zwPVMhG9W8UjwvqxHXfdlTH37nThsN\n-----END RSA PRIVATE KEY-----")
}

func createLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}
