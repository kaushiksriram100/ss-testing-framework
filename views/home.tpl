<h1>Stream Splitter Plugins Testing Framework</h1>
<form action="/process" method="post">
<h4 align="left">Input Data (please limit to 50 events approx per test) and stream-splitter plugins.conf (only contents inside plugins block. exclude string literal 'plugins {'</h4>
<textarea style="width:600px; height:320px;" rows="15" cols="50" name="inputtext" align="left"></textarea>
<textarea style="width:400px; height:320px;" rows="10" cols="50" name="applicationconf" align="right"></textarea>
<br/>
<br/>
<button type="submit" class="btn">Submit</button>
<br/>
<h4>Results:</h4>
<textarea style="width:700px; height:320px;" readonly >{{.OUTPUT}}</textarea>


{{if .ERROR}}<h4 style="color:red"> errors:  </h4> {{.ERROR}} {{else}} <h4 style="color:green"> No Errors </h4>{{end}}
<br />
<br />
<ul>
<li>If error is no data in output file then it _is_ possible that input stream didn't match the plugins and translated to nil output.</li>
</ul>
<br />
</form>
<h4> Sample plugins.conf for hubble_stream data</h4>
<textarea style="width:700px; height:320px;" readonly >
   map-plugin {
     name = "JSON-map-plugin"
     consumer = [
       ${stream-splitter.consumer}
     ]
   }

                HTTPLOG_error_plugin {
      name = "filter-plugin"
      rename-to = "httplog_errors"
      filters {
        operator = AND
        operands = [
          {
            expression = "request->requestUrl"
            values = ["/rum.gif"]
          }
          ,
          {
            expression = "request->queryStringMap->ctx"
            values = ["ErrorPage"]
          },
          {
            expression = "request->queryStringMap->a"
            values = ["ERRORPAGE_VIEW"]
          }
        ]

      }
      consumer = [
        ${stream-splitter.plugins.map-plugin}
      ]
      producers = [${stream-splitter.producers.file}]
    }

</textarea>


<h5>*******************DEV NOTES********************************************</h5>

<ul>
<li>Author: Sriram Kaushik</li>
<li>Team: Telemetry Platform</li>
<li>Use output (producer) file as `/var/tmp/testing-output.txt` and consumer file as `/var/tmp/testing-input.txt`
<li>Confirm if the appropriate SS version and Topology is deployed in this QA env via backend ansible.</li>
<li> Use this project for deployment: https://gecgithub01.walmart.com/pulse/ansible-stream-splitter/tree/master/inventories/testing/CDC/prod</li>
</ul>
<br />