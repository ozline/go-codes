
a:
	echo "a"

SERVICES := api user follow interaction video chat
service = $(word 1, $@)
endpoint = $(word 2, $@)

.PHONY: $(SERVICES)
$(SERVICES):
	echo $(endpoint)