<h1>Stream Splitter Plugins Testing Framework</h1>
<form action="/process" method="post">
<h4>Input Data (please limit to 60 events approx per test)</h4>
<textarea rows="15" cols="50" name="inputtext"></textarea>
<br/>
<br/>
<button type="submit" class="btn">Submit</button>
<br/>
<h4>Results:</h4>
<textarea readonly>{{.OUTPUT}}</textarea>


{{if .ERROR}}<h4> errors:  </h4> {{.ERROR}} {{else}} <h4> No Errors </h4>{{end}}
</form>