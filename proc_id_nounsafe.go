// +build fastrand_nounsafe

package fastrand

const fastrand_nounsafe = true

func procPin() int { return -1 }

func procUnpin() {}
