include $(GOROOT)/src/Make.$(GOARCH)

TARG=scheme
GOFILES=\
	scheme.go

include $(GOROOT)/src/Make.pkg

format:
	echo $(GOFILES) | xargs gofmt -w

man: man/scheme.3

man/scheme.3:
	ron -b man/scheme.3.ron
