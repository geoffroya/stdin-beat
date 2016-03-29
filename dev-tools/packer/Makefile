DATE:=$(shell date +%y%m%d%H%M%S)
BUILDID?=$(DATE)

.PHONY: all
all: stdin-beat/deb stdin-beat/rpm stdin-beat/darwin stdin-beat/win stdin-beat/bin build/upload/build_id.txt latest

ES_BEATS=../../vendor/github.com/elastic/beats
include $(ES_BEATS)/dev-tools/packer/scripts/Makefile


.PHONY: stdin-beat
stdin-beat: build/upload
	# cross compile on ubuntu
	cd build && xgo -image=tudorg/beats-builder \
		-before-build=../$(ES_BEATS)/dev-tools/packer/xgo-scripts/before_build.sh \
		-pkg $@ \
		-source ../../.. \
		src/github.com/geoffroya/stdin-beat/$@
	# linux builds on debian 6
	cd build && xgo -image=tudorg/beats-builder-deb6 \
		-before-build=../$(ES_BEATS)/dev-tools/packer/xgo-scripts/before_build.sh \
		-pkg $@ \
		-source ../../.. \
		src/github.com/geoffroya/stdin-beat/$@

