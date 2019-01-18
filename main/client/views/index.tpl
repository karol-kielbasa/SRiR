<!DOCTYPE html>

<html>
    <head>
        <title>Beego</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
         <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap-theme.min.css">
    </head>
    <body>
        <header class="hero-unit">
            <div class="container">
                <div class="row">
                  <div class="hero-text">
                    <h2>Select file</h2>
                    <form action="/sendFile" method="post" enctype="multipart/form-data">
                        <input id="input" type="file" name="file" multiple="multiple">
                        <input type="submit" name="submit" value="Send file">
                    </form>
                    <div id="build">
                        <h2>Build selected file</h2>
                        <form action="/buildFile" method="get" >
                            <input id="query-param-build" type="hidden" name="fileName"/>
                            <input type="submit" onclick="getFileName()">
                        </form>
                    </div>
                    <div id="execute">
                        <h2> Execute selected file </h2>
                        <form action="/executeFile" method="get">
                            <input id="query-param-execute" type="hidden" name="fileName"/>
                            <input type="submit" onclick="getFileName()">
                        </form>
                    </div>
                    <div id="backup">
                        <h2> Backup file and get comparison reports </h2>
                        <form action="/backupFile" method="get">
                            <input id="query-param-backup" type="hidden" name="fileName"/>
                            <input type="submit" onclick="getFileName()">
                        </form>
                    </div>
                  </div>
                </div>
            </div>
        </header>
        <script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
        <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
    </body>
</html>
<script>
  function getFileName(){
      let val = $('#input').val().replace(/C:\\fakepath\\/i, '')
      if(val !== ""){
         $( "#query-param-execute" ).val(val)
         $( "#query-param-build" ).val(val)
         $( "#query-param-backup" ).val(val)
      }
    }
</script>