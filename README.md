# Stats-Go - Yet another library for metrics instrumentation.
[![Build Status](https://travis-ci.com/zviadm/stats-go.svg?branch=master)](https://travis-ci.com/zviadm/stats-go)

In spirit, this library has similar goals as [opencensus-go](https://github.com/census-instrumentation/opencensus-go) but
with way less features and way less code overall. This library is most similar to
[client_golang/prometheus](https://github.com/prometheus/client_golang), but with a different approach on stats collection.

Stats-Go is made up of few separate GO modules:
- [metrics](./metrics): Provides library for instrumenting GO code. Performance is a big priority for this library
to allow instrumenting even hot code paths without much worry about additional cpu utilization or memory allocations.
APIs for instrumenting the code are not expected to change in backwards incompatible ways, even in experimental phase.
- [exporters](./exporters): Provides GO modules for exporting collected metrics to different potential backends.
- [handlers](./handlers): Provides various separate libraries to help instrument core common libraries. These libraries
are in iteration phase, thus they might change pretty substantially.
