const conf = require('./gulp.conf');

module.exports = function (config) {
  const configuration = {
    basePath: '../',
    singleRun: false,
    autoWatch: true,
    logLevel: 'INFO',
    junitReporter: {
      outputDir: 'test-reports'
    },
    browsers: [
      'PhantomJS'
    ],
    frameworks: [
      'jasmine',
      'jspm'
    ],
    jspm: {
      loadFiles: [
        conf.path.src('app/**/*.js')
      ],
      config: 'jspm.config.js',
      browser: 'jspm.test.js'
    },
    plugins: [
      require('karma-jasmine'),
      require('karma-junit-reporter'),
      require('karma-coverage'),
      require('karma-phantomjs-launcher'),
      require('karma-phantomjs-shim'),
      require('karma-jspm')
    ]
  };

  config.set(configuration);
};
