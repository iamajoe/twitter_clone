import { AChart } from "./components/chart.js";
import { Test } from "./components/test.js";

const componentMap = {
  "app-test": Test,
  "app-chart": AChart,
};

const componentKeys = Object.keys(componentMap);
for (let i = 0; i < componentKeys.length; i += 1) {
  const k = componentKeys[i];
  window.customElements.define(k, componentMap[k]);
}
