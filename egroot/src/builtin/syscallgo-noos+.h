// +build noos

#define GO(call, lock) do {                   \
	void func() {                             \
		call;                                 \
		builtin$Syscall1(builtin$DELTASK, 0); \
	}                                         \
	builtin$NewTask(func, lock);              \
} while(0)

static inline
void goready() {
	builtin$Syscall0(builtin$TASKUNLOCK);
}

static inline
uintptr builtin$b2u(bool b) {
	return (uintptr)(b);
}

static inline
uintptr builtin$f2u(void (*f)()) {
	union {void (*in)(); uintptr out;} cast;
	cast.in = f;
	return cast.out;
}

/*
static
void newTask(void (*func)(), bool lock) {
	uintptr$$uintptr r = builtin$Syscall2(
		builtin$NEWTASK, (uintptr)(func), lock
	);
	if (r._1 != 0) {
			panic(INTERFACE(r._1, nil));
	}
}
*/
