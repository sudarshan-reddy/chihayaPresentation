// START OMIT
_, err := conn.Do("ZREMRANGEBYSCORE", Key, "-inf", fmt.Sprintf("(%d", time.Now().Add(-s.peerLifetime).Unix()))
// END OMIT
