const { defineConfig } = require("cypress");

module.exports = defineConfig({
  projectId: "qtbmb8",

  component: {
    devServer: {
      framework: "next",
      bundler: "webpack",
    },
  },

  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
