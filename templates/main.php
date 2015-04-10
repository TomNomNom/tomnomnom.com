<!doctype html>
<html lang="en">
  <head>
    <title><?php echo $this->e($title)?: "Tom Hudson's blog - TomNomNom.com"; ?></title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <link href="http://fonts.googleapis.com/css?family=Droid+Serif" rel="stylesheet" type="text/css">
    <link href="http://fonts.googleapis.com/css?family=Droid+Sans" rel="stylesheet" type="text/css">
    <link href="http://fonts.googleapis.com/css?family=Droid+Sans+Mono" rel="stylesheet" type="text/css">
    <link href="/styles/main.css" rel="stylesheet" type="text/css"/>
    <link href="/styles/code-dark.css" rel="stylesheet" type="text/css"/>
  </head>
  <body>
    <header>
      <h1><a href="/">TomNomNom.com</a></h1>
      <h2>Blogging <strike>since 2008</strike> until 2010</h2>
    </header>

    <section class="main">
      <?php $this->render($subpage, $subpageData); ?>
    </section>

    <footer>
      &copy; Tom Hudson 2015
    </footer>
  </body>

  <?php if(isSet($gaTracking) && $gaTracking): ?>
    <script type="text/javascript">
      var _gaq = _gaq || [];
      _gaq.push(['_setAccount', 'UA-22278243-3']);
      _gaq.push(['_trackPageview']);

      (function() {
        var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
        var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
      })();
    </script>
  <?php endif; ?>
</html>
