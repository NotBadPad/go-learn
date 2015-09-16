package main

type I interface {
	DoSomeWork()
}

type T struct {
	a int
}

func (t *T) DoSomeWork() {
}

func main() {
	t := &T{}
	i := I(t)
	print(i)
}

/**
 go tool 6g -W ./test1.go


before (*T).DoSomeWork
after walk (*T).DoSomeWork

before main
.   DCL l(15)
.   .   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T

.   AS l(15) colas(1) tc(1)
.   .   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T
.   .   PTRLIT l(15) esc(no) ld(1) tc(1) PTR64-*main.T
.   .   .   STRUCTLIT l(15) tc(1) main.T
.   .   .   .   TYPE <S> l(15) tc(1) implicit(1) type=PTR64-*main.T PTR64-*main.T

.   DCL l(16)
.   .   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I

.   AS l(16) tc(1)
.   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T
.   .   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T

.   AS l(16) colas(1) tc(1)
.   .   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I
.   .   CONVIFACE l(16) tc(1) main.I
.   .   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

.   VARKILL l(16) tc(1)
.   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

.   PRINT l(17) tc(1)
.   PRINT-list
.   .   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I
after walk main
.   DCL l(15)
.   .   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T

.   AS-init
.   .   AS l(15) tc(1)
.   .   .   NAME-main.autotmp_0002 u(1) a(1) l(15) x(0+0) class(PAUTO) esc(N) tc(1) used(1) main.T

.   .   AS l(15) tc(1)
.   .   .   NAME-main.autotmp_0001 u(1) a(1) l(15) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T
.   .   .   ADDR u(2) l(15) tc(1) PTR64-*main.T
.   .   .   .   NAME-main.autotmp_0002 u(1) a(1) l(15) x(0+0) class(PAUTO) esc(N) tc(1) used(1) main.T

.   .   AS u(2) l(15) tc(1)
.   .   .   IND u(2) l(15) tc(1) main.T
.   .   .   .   NAME-main.autotmp_0001 u(1) a(1) l(15) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T
.   AS u(100) l(15) tc(1)
.   .   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T
.   .   NAME-main.autotmp_0001 u(1) a(1) l(15) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

.   DCL l(16)
.   .   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I

.   AS u(2) l(16) tc(1)
.   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T
.   .   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T

.   AS-init
.   .   AS l(16) tc(1)
.   .   .   NAME-main.autotmp_0003 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*uint8
.   .   .   NAME-go.itab.*"".T."".I u(1) a(1) l(16) x(0+0) class(PEXTERN) tc(1) used(1) PTR64-*uint8

.   .   IF l(16) tc(1)
.   .   IF-test
.   .   .   EQ l(16) tc(1) bool
.   .   .   .   NAME-main.autotmp_0003 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*uint8
.   .   .   .   LITERAL-nil u(1) a(1) l(16) tc(1) PTR64-*uint8
.   .   IF-body
.   .   .   AS l(16) tc(1)
.   .   .   .   NAME-main.autotmp_0003 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*uint8
.   .   .   .   CALLFUNC u(100) l(16) tc(1) PTR64-*byte
.   .   .   .   .   NAME-runtime.typ2Itab u(1) a(1) l(2) x(0+0) class(PFUNC) tc(1) used(1) FUNC-funcSTRUCT-(FIELD-
.   .   .   .   .   .   NAME-runtime.typ·2 u(1) a(1) l(2) x(0+0) class(PPARAM) f(1) PTR64-*byte PTR64-*byte, FIELD-
.   .   .   .   .   .   NAME-runtime.typ2·3 u(1) a(1) l(2) x(8+0) class(PPARAM) f(1) PTR64-*byte PTR64-*byte, FIELD-
.   .   .   .   .   .   NAME-runtime.cache·4 u(1) a(1) l(2) x(16+0) class(PPARAM) f(1) PTR64-*PTR64-*byte PTR64-*PTR64-*byte) PTR64-*byte
.   .   .   .   CALLFUNC-list
.   .   .   .   .   AS u(2) l(16) tc(1)
.   .   .   .   .   .   INDREG-SP a(1) l(16) x(0+0) tc(1) runtime.typ·2 G0 PTR64-*byte
.   .   .   .   .   .   ADDR u(2) a(1) l(16) tc(1) PTR64-*uint8
.   .   .   .   .   .   .   NAME-type.*"".T u(1) a(1) l(11) x(0+0) class(PEXTERN) tc(1) uint8

.   .   .   .   .   AS u(2) l(16) tc(1)
.   .   .   .   .   .   INDREG-SP a(1) l(16) x(8+0) tc(1) runtime.typ2·3 G0 PTR64-*byte
.   .   .   .   .   .   ADDR u(2) a(1) l(16) tc(1) PTR64-*uint8
.   .   .   .   .   .   .   NAME-type."".I u(1) a(1) l(16) x(0+0) class(PEXTERN) tc(1) uint8

.   .   .   .   .   AS u(2) l(16) tc(1)
.   .   .   .   .   .   INDREG-SP a(1) l(16) x(16+0) tc(1) runtime.cache·4 G0 PTR64-*PTR64-*byte
.   .   .   .   .   .   ADDR u(2) a(1) l(16) tc(1) PTR64-*PTR64-*uint8
.   .   .   .   .   .   .   NAME-go.itab.*"".T."".I u(1) a(1) l(16) x(0+0) class(PEXTERN) tc(1) used(1) PTR64-*uint8
.   AS u(100) l(16) tc(1)
.   .   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I
.   .   EFACE u(2) l(16) tc(1) main.I
.   .   .   NAME-main.autotmp_0003 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*uint8
.   .   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

.   VARKILL l(16) tc(1)
.   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

.   EMPTY-init
.   .   CALLFUNC u(100) l(17) tc(1)
.   .   .   NAME-runtime.printiface u(1) a(1) l(2) x(0+0) class(PFUNC) tc(1) used(1) FUNC-funcSTRUCT-(FIELD-main.I)
.   .   CALLFUNC-list
.   .   .   AS u(1) l(17) tc(1)
.   .   .   .   INDREG-SP a(1) l(17) x(0+0) tc(1) main.I
.   .   .   .   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I
.   EMPTY u(100) l(17) tc(1)

before init
.   IF l(18) tc(1)
.   IF-test
.   .   NE l(18) tc(1) bool
.   .   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   .   LITERAL-0 u(1) a(1) l(18) tc(1) uint8
.   IF-body
.   .   IF l(18) tc(1)
.   .   IF-test
.   .   .   EQ l(18) tc(1) bool
.   .   .   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   .   .   LITERAL-2 u(1) a(1) l(18) tc(1) uint8
.   .   IF-body
.   .   .   RETURN l(18) tc(1)

.   .   CALLFUNC l(18) tc(1)
.   .   .   NAME-runtime.throwinit u(1) a(1) l(2) x(0+0) class(PFUNC) tc(1) used(1) FUNC-funcSTRUCT-()

.   AS l(18) tc(1)
.   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   LITERAL-1 u(1) a(1) l(18) tc(1) uint8

.   AS l(18) tc(1)
.   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   LITERAL-2 u(1) a(1) l(18) tc(1) uint8

.   RETURN l(18) tc(1)
after walk init
.   IF l(18) tc(1)
.   IF-test
.   .   NE u(2) l(18) tc(1) bool
.   .   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   .   LITERAL-0 u(1) a(1) l(18) tc(1) uint8
.   IF-body
.   .   IF l(18) tc(1)
.   .   IF-test
.   .   .   EQ u(2) l(18) tc(1) bool
.   .   .   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   .   .   LITERAL-2 u(1) a(1) l(18) tc(1) uint8
.   .   IF-body
.   .   .   RETURN l(18) tc(1)

.   .   CALLFUNC u(100) l(18) tc(1)
.   .   .   NAME-runtime.throwinit u(1) a(1) l(2) x(0+0) class(PFUNC) tc(1) used(1) FUNC-funcSTRUCT-()

.   AS u(2) l(18) tc(1)
.   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   LITERAL-1 u(1) a(1) l(18) tc(1) uint8

.   AS u(2) l(18) tc(1)
.   .   NAME-main.initdone· u(1) a(1) l(18) x(0+0) class(PEXTERN) tc(1) used(1) uint8
.   .   LITERAL-2 u(1) a(1) l(18) tc(1) uint8

.   RETURN l(18) tc(1)

before I.DoSomeWork
.   CALLINTER l(20) tc(1)
.   .   DOTINTER l(20) x(0+0) tc(1) FUNC-methodSTRUCT-(FIELD-PTR64-*STRUCT-struct {}) funcSTRUCT-()
.   .   .   NAME-main..this u(1) a(1) g(1) l(20) x(0+0) class(PPARAM) f(1) tc(1) used(1) main.I
.   .   .   NAME-main.DoSomeWork u(1) a(1) l(20) x(0+0)
after walk I.DoSomeWork
.   CALLINTER u(100) l(20) tc(1)
.   .   DOTINTER u(2) l(20) x(0+0) tc(1) FUNC-methodSTRUCT-(FIELD-PTR64-*STRUCT-struct {}) funcSTRUCT-()
.   .   .   NAME-main..this u(1) a(1) g(1) l(20) x(0+0) class(PPARAM) f(1) tc(1) used(1) main.I
.   .   .   NAME-main.DoSomeWork u(1) a(1) l(20) x(0+0)




*/
