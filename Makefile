SUBDIRS := $(wildcard */.)

all: $(SUBDIRS)
$(SUBDIRS):
	cd $@ && go install

.PHONY: all $(SUBDIRS)