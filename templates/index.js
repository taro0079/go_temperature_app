< !DOCTYPE html >
	<html lang="ja">
		<head>
			<meta charset="UTF-8">
			<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.5.0/Chart.js"></script>
			<title>Sample Application</title>

		</head>
		<body>
			<h1>コロナワクチン副反応　体温推移</h1>
			<form method="post" action="/new">
				<p>Time<input type="datetime-local" name="time" placeholder="add"></input></p>
				<p>Temperature<input type="number" name="temp" placeholder="add"></input></p>

				<p><input type="submit" value="Send"></input></p>
			</form>

			<div style="height:350px; width:500px; margin:0 auto;">
				<canvas id="ChartID"></canvas>
			</div>

			<ul>
				{{ range .measurement }}
				<li>Time：{{ .Time }},
					<ul>
						<li>
							Temperature：{{ .Temperature }}
						</li>
					</ul>
					<label><a href="/detail/{{.ID}}">編集</a></label>
					<label><a href="/delete_check/{{.ID}}">削除</a></label>
				</li>
				<br></br>
				{{ end }}
			</ul>
			<script>
				var ctx = document.getElementById('ChartID').getContext('2d');
				var myChart = new Chart(ctx, {
					type: タイプ,
					data: データ,
					options: オプション});
			</script>
		</body>
	</html>




