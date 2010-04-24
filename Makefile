# prerequisite: GOROOT and GOARCH must be defined

# defines $(GC) (compiler), $(LD) (linker) and $(O) (architecture)
include $(GOROOT)/src/Make.$(GOARCH)

# name of the package (library) being built
TARG=vtgostack

GOFILES=\
	stack.go\

include $(GOROOT)/src/Make.pkg