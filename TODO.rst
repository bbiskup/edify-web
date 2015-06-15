TODOs
=====

Bugs & Shortcomings
-------------------

- Support UNH, UNT etc. or give warning message about unparsed data elements
- display message spec with groups (currently only linear list with only 
  top-level (trigger) segments)
- bug? language name code (3453) missing from UNCL spec. 
- check if code specs of simple data elem are nil and display appropriate msg
  example: http://localhost:8001/specs/simpledataelement/6066
  

Features
--------

- Tests & benchmarks
  - use httptest.HttpResponseRecorder?
  - Dependency injection à la http://www.markjberger.com/testing-web-apps-in-golang/ ?
- Sort code search results
- Codes list browseable/searchable
- Expose additional spec info (notes, full description)
- File upload of spec
- REST API for validating messages?
- edify: give list of available (downloaded/extracted) spec versions
- Spec search: fill combo box with available versions (default: most reccent)
- structured output of parsed message (Dynatree?)
- Display of (original?) message with tooltips containing spec information?
- Rendering of messages with template
  - future (maybe): template as form, to allow creation of message
- store uploaded messages in DB?
- import EDIFACT messags from OpenData source?
- User account?
- Ansible scripts for public server
- Parallel search for specs?
  - Check if this could really be faster; BenchmarkSpecSearchNoResults
  - possibly: parallel rendering of templates?
- List of all specs of a kind at URL /specs/message/, /specs/segment/, ....
- Spec search: support data elements
- Spec search: support codes
- Spec search: organize different kinds of specs in accordion or tabs
  - display badge with result count
  - omit spec types with no results?
- Spec search. (maybe) limit search to particular kinds of specs
- Pagination of long tables (example: long code list http://localhost:8001/specs/simpledataelement/4451)
  - client-side pagination for simplicity
  - link at bottom of page to jump to top of page



