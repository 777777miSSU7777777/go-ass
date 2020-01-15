const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = ( env, argv ) => {
	return ({
		entry: "./src/index.jsx",
		output: {
			path: path.join(__dirname, "/dist"),
			filename: "index-bundle.js",
			publicPath: argv.mode === "production" ? "/static/" : ""
		},
		module: {
			rules: [
				{
					test: /\.(js|jsx)$/,
					exclude: /node_modules/,
					use: ["babel-loader"]
				},
				{
					test: /\.css$/,
					use: ["style-loader", "css-loader"]
				},
				{
					test: /\.(png|jpg)$/,
					loader: "url-loader"
				},
			]
		},
		externals: {
			Config: JSON.stringify(argv.mode === "production" ? require("./config.prod.json") : require("./config.dev.json"))
		},
		plugins: [
			new HtmlWebpackPlugin({
				template: "./src/index.html"
			})
		]
	});
}