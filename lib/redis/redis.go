package redis

import (
	"fmt"
	"log"
	"time"

	"todo/utils/errors"

	redisClient "github.com/go-redis/redis"
	"github.com/go-redsync/redsync/v3"
	"github.com/go-redsync/redsync/v3/redis"
	"github.com/go-redsync/redsync/v3/redis/goredis"
)

const (
	RedisTopupPrefix              = "top_up:"
	RedisInvoicePrefix            = "invoice:"
	RedisBnibProductPrefix        = "bnib_product:"
	RedisBnibDirectBuyOrderPrefix = "direct_bnib_buy_order:"
	RedisBnibDirectProductPrefix  = "direct_bnib_product:"
	RedisBnibBuyOrderPrefix       = "bnib_buy_order:"
	RedisBnibProductPricePrefix   = "bnib_product_price:"
	RedisBnibBuyOrderPricePrefix  = "bnib_buy_order_price:"
	RedisBnibTransactionPrefix    = "bnib_transaction:"
	RedisLock                     = "lock:"
	RedisSellingHistoryPrefix     = "selling_history:"
	// mercury
	RedisRetailPrefix        = "retail:"
	RedisStoreClosedCooldown = "store_closed_cooldown:"
	// covid
	RedisFolderNamePrefix        = "folder_name:"
	RedisAuthenticationPrefix    = "authentication:"
	RedisPackageSetPrefix        = "package_set:"
	RedisLegitCheckInvoicePrefix = "legit_check_invoice:"
	RedisPackageInvoicePrefix    = "package_invoice:"
	RedisPaymentGatewayPrefix    = "payment_gateway:"

	RedisExceptionLoggingHandlePrefix = "exception_logging_handle:"

	RedisQueryCachePrefix = "query_cache:"
)

type ExpirationTimeName string

type ExpirationTime struct {
	Name     string    `json:"name"`
	Duration time.Time `json:"duration"`
}

const (
	RedisExpirationTimeOneMinute     ExpirationTimeName = "one_minute"
	RedisExpirationTimeFiveMinutes   ExpirationTimeName = "five_minute"
	RedisExpirationTimeTenMinutes    ExpirationTimeName = "ten_minute"
	RedisExpirationTimeThirtyMinutes ExpirationTimeName = "thirty_minute"
	RedisExpirationTimeOneHour       ExpirationTimeName = "one_hour"
	RedisExpirationTimeSixHour       ExpirationTimeName = "six_hour"
	RedisExpirationTimeOneDay        ExpirationTimeName = "one_day"
	RedisExpirationTimeOneWeek       ExpirationTimeName = "one_week"
)

func (e ExpirationTimeName) ToTimeDuration() time.Duration {
	mappings := map[ExpirationTimeName]time.Duration{
		RedisExpirationTimeOneMinute:     time.Minute,
		RedisExpirationTimeFiveMinutes:   5 * time.Minute,
		RedisExpirationTimeTenMinutes:    10 * time.Minute,
		RedisExpirationTimeThirtyMinutes: 30 * time.Minute,
		RedisExpirationTimeOneHour:       time.Hour,
		RedisExpirationTimeSixHour:       6 * time.Hour,
		RedisExpirationTimeOneDay:        24 * time.Hour,
		RedisExpirationTimeOneWeek:       7 * 24 * time.Hour,
	}
	return mappings[e]
}

type Credentials struct {
	Host     string
	Port     string
	Password string
}

type Client interface {
	Get(prefix string, key string) string
	Set(prefix string, key string, value string, expirationTime time.Duration) error
	Delete(prefix string, key string) error
	DeleteMatch(pattern string) error
	Keys(pattern string) []string
	Ping() error
	Close() error
	NewMutex(key string) *redsync.Mutex
	SAdd(prefix string, key string, members interface{}) error
	SRem(prefix string, key string, members interface{}) error
	SMembers(prefix string, key string) ([]string, error)
	//CreateLock(key string) error
	//CheckLock(key string) (bool, error)
	//ReleaseLock(key string) error
}

type Redis struct {
	client  *redisClient.Client
	redsync *redsync.Redsync
}

func NewClient(credentials Credentials, appEnv string) Client {
	client := redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:%s", credentials.Host, credentials.Port),
		Password: credentials.Password,
		DB:       0,
	})
	status := client.Ping()
	if status.Err() != nil {
		if appEnv != "development" {
			log.Panic(status.Err())
		} else {
			log.Println("warning: redis not connected")
		}
	}

	pool := goredis.NewGoredisPool(client)
	rs := redsync.New([]redis.Pool{pool})

	return &Redis{
		client:  client,
		redsync: rs,
	}
}

func GetExpirationTimeDurationFromName(name string) time.Duration {
	expirationTimeNames := []ExpirationTimeName{
		RedisExpirationTimeOneMinute,
		RedisExpirationTimeFiveMinutes,
		RedisExpirationTimeTenMinutes,
		RedisExpirationTimeThirtyMinutes,
		RedisExpirationTimeOneHour,
		RedisExpirationTimeSixHour,
		RedisExpirationTimeOneDay,
		RedisExpirationTimeOneWeek,
	}

	for _, expirationTimeName := range expirationTimeNames {
		if ExpirationTimeName(name) == expirationTimeName {
			return expirationTimeName.ToTimeDuration()
		}
	}
	return 0
}

func (r *Redis) Keys(pattern string) []string {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return nil
	}

	val, err := r.client.Keys(pattern).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			fmt.Println(err)
		}
	}
	return val
}

func (r *Redis) Get(prefix string, key string) string {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return ""
	}

	val, err := r.client.Get(fmt.Sprint(prefix, key)).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			fmt.Println(err)
		}
	}
	return val
}

func (r *Redis) Set(prefix string, key string, value string, expirationTime time.Duration) error {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return errors.ErrUnprocessableEntity
	}

	err := r.client.Set(fmt.Sprint(prefix, key), value, expirationTime).Err()
	if err != nil {
		fmt.Println(err)
		return errors.ErrUnprocessableEntity
	}

	return nil
}

func (r *Redis) Delete(prefix string, key string) error {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return errors.ErrUnprocessableEntity
	}

	err := r.client.Del(fmt.Sprint(prefix, key)).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (r *Redis) DeleteMatch(pattern string) error {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return errors.ErrUnprocessableEntity
	}

	keys := r.Keys(fmt.Sprintf("*%s*", pattern))
	for _, key := range keys {
		err := r.client.Del(key).Err()
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	return nil
}

func (r *Redis) Ping() error {
	pong, err := r.client.Ping().Result()
	if err != nil {
		log.Println("error pinging redis:", err)
		return errors.ErrUnprocessableEntity

	}

	fmt.Println("connected to redis:", pong)
	return nil
}

func (r *Redis) Close() error {
	err := r.client.Close()
	if err != nil {
		log.Println("error closing redis:", err)
		return err
	}
	return nil
}

func (r *Redis) NewMutex(key string) *redsync.Mutex {
	return r.redsync.NewMutex(key)
}

// Add one or more members to a set
func (r *Redis) SAdd(prefix string, key string, members interface{}) error {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return errors.ErrUnprocessableEntity
	}

	err := r.client.SAdd(fmt.Sprint(prefix, key), members).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Remove one or more members from a set
func (r *Redis) SRem(prefix string, key string, members interface{}) error {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return errors.ErrUnprocessableEntity
	}

	err := r.client.SRem(fmt.Sprint(prefix, key), members).Err()
	if err != nil {
		fmt.Println(err)
		return errors.ErrUnprocessableEntity
	}
	return err
}

// Get all the members in a set
func (r *Redis) SMembers(prefix string, key string) ([]string, error) {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return nil, errors.ErrUnprocessableEntity
	}

	members, err := r.client.SMembers(fmt.Sprint(prefix, key)).Result()
	if err != nil {
		fmt.Println(err)
		return nil, errors.ErrUnprocessableEntity
	}
	return members, nil
}

// Set a key's time to live in seconds
func (r *Redis) Expire(key string, expiration time.Duration) error {
	if r.client == nil {
		log.Println("warning: redis not connected")
		return errors.ErrUnprocessableEntity
	}

	err := r.client.Expire(key, expiration).Err()
	if err != nil {
		fmt.Println(err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
