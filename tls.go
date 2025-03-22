package rsa

func main() {

	p := 11
	q := 17

	e := 3

	n := p * q
	phi_n := (p - 1) * (q - 1)
	d := egcd(e, phi_n)

}

func gcd(m, n int) int {
	if (n  == 0) {
		return m;

	}
	return gcd(n, m % n)
}

func egcd(a, b int) (int, int, int) {
	x, y, u, v := 0, 1, 1, 0
	while a != 0
	{
		q, r := int(b/a), b % a
		m, n := x - u * q, y - v*q
		b,a x,y, u,v = a,r, u,v, m,n
	}
	gcd = b
	return gcd, x, v
}

func modInv(a, phi int) int {
	gcd, x, y = egcd(a, phi)
	if gcd != 1 {
		return -1
	}else {
		return x % phi
	}
}
