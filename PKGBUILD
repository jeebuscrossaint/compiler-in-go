# Maintainer: 6ty apatel6ty@protonmail.com

pkgname=constrict
pkgver=1.0.0
pkgrel=1
pkgdesc="A simpler solution to make, written for only myself but welcome to the public."
arch=('x86_64')
url=https://jeebuscrossaint.github.io
license=('MIT')
depends=('go')
source = add source here
sha256sums= add sha256 here

build() {
	cd "$srcdir/$pkgname-$pkgver"
	./configure --prefix=/usr
	go build src/main.go
	mv main constrict
}

package() {
	cd "$srcdir/$pkgname-$pkgver"
	make DESTDIR=$pkgdir/" install
}
