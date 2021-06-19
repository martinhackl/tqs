PACKAGE_NAME=tqs

BUILD_DIR=build/
BUILD_FILE=$(BUILD_DIR)$(PACKAGE_NAME)

make: build run

build:
	go build -o $(BUILD_FILE)

clean:
	rm -f $(BUILD_FILE)

run:
	$(BUILD_FILE)