BUILD_DIR=bin
INSTALL_DIR=~/bin
EXECUTABLE=gotxt

.DEFAULT_GOAL := bin
.PHONY: cleanall bin

bin: cleanall
	go build
	mkdir $(BUILD_DIR)
	mv $(EXECUTABLE) $(BUILD_DIR)

cleanall:
	rm -f $(EXECUTABLE)
	rm -rf $(BUILD_DIR)

install: bin
	cp $(BUILD_DIR)/$(EXECUTABLE) $(INSTALL_DIR)/$(EXECUTABLE)
