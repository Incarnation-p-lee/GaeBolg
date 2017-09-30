src     :=$(shell ls *.go)

TARGET  :=manual
CC      :=go build
CLEAN   :=go clean
INTRP   :=go run

.PHONY: $(TARGET) clean intrp

$(TARGET):$(src)
	$(CC) -o $@ $^

clean:
	$(CLEAN)

intrp:$(src)
	$(INTRP) $^

