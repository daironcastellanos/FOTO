import { defineConfig } from "cypress";

export default defineConfig({
  projectId: 'qtbmb8',
  component: {
    devServer: {
      framework: "next",
      bundler: "webpack",
    },
  },
});
