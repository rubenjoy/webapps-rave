# Web Frameworks Benchmarks

Let bring all webapp framework to a test-bench with load-generator such as:
_Apache JMeter_, _Locust.io_, etc. Anyway Techempower already did the
[benchmark](https://www.techempower.com/benchmarks/) testing a full-stack
application. To differentiate, this benchmark will focus on web framework
which many backend developers use it as racehorse to serve JSON REST API.
Considering that, all tests here are should be containerized, repeatable, and
having easy-to-grasp report.

## Ground Rule in this Performance Field

Before starting the test, we need a guideline first.

### REST API WebApp

 - all response should be in JSON.
 - it serve healtcheck response, containing: ok-message and server time
 - it serve ping request that send pong message
 - it serve fortune cookie with Unix-fortune format, '%' as separator

### Container Resource Allocation

Before watching Henning Jacobs'
[video](https://www.youtube.com/watch?v=eBChCFD9hfs), I never realized
that resource constraint could cause unforeseeable latency. Just curious,
how to create those effects in this test so that I can learn more about it.
For now, I am considering only memory and CPU constraint and not yet deciding
on resource declaration.

### Load Generator Tool

I opt two tools: _Apache JMeter_ and _Locust.io_, which I am familiar with.
The tools used here should comply with the following requirements:

 - have ability to make multiple connection, not having single reusable
   connection
 - store all samples of response measurement in CSV format, so that the report
   could be generated anytime
 - able to run in headless mode and containerized
 - extendable or _hackable_

### Report

 - average, minimum, maximum, and 99-percentile response time
 - should be in easy-to-published format: HTML or PDF
 - contains information about resource consumption, scalability
 - all resource limitation implied in container, if any, should be reported
 - any OOM or restart event inside container should be reported too

## Roadmap

The tests aim for programming language agnostic, but for now only listed languages and frameworks are on the test.

- Golang
  - net/http (Go package)
  - gorilla/mux
  - revel/revel
  - labstack/echo
  - gin-gonic/gin
  - [valyala/fasthttp](https://github.com/valyala/fasthttp) with plain vanilla HandlerFunc
- Node JS
  - Express
  - Hapijs
  - http.Server (Node JS API)
 - Python
  - Django
  - Flask
  - WSGI (python vanilla)
 - Groovy
  - Grails
  - Servlet
  - Ratpack
  - javaConductor/gServ
  - Spring Framework
  - Play Framework
 - Functional Programming Language
  - Scala
  - Haskell

## Beyond the Future Roadmap

 - Add APM to have better understanding about what happen behind the scene
 - Pivot the bench to security-test