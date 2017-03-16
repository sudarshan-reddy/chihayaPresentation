// START ZREMRANGEBYSCORE
_, err := conn.Do("ZREMRANGEBYSCORE", Key, "-inf", fmt.Sprintf("(%d", time.Now().Add(-s.peerLifetime).Unix()))
// END ZREMRANGEBYSCORE

// START EXPIRE
conn.Send("EXPIRE", Key, int(s.peerLifetime.Seconds()))
// END EXPIRE

