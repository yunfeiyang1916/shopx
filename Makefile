
.PHONY: api
# generate api
api:
	find app -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && make api'

.PHONY: wire
# generate wire
wire:
	find app -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && make wire'

.PHONY: proto
# generate proto
proto:
	find app -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && make proto'