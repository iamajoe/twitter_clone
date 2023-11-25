export class AChart extends HTMLElement {
  _chart = undefined;

  async getLib() {
    const cdn = "https://cdn.jsdelivr.net/npm/chart.js";
    if (window.Chart == null) {
      await import(cdn);
    }

    return window.Chart;
  }

  parseData() {
    const dataRaw = this.getAttribute("data");
    let data = [];
    if (dataRaw != null && dataRaw.length > 0) {
      data = JSON.parse(dataRaw);
      if (!Array.isArray) {
        data = [data];
      }
    }

    return data;
  }

  // Custom element added to page
  async connectedCallback() {
    const shadowRoot = this.attachShadow({ mode: "open" });
    shadowRoot.innerHTML = '<canvas class="chart"/>';
    const ctx = shadowRoot.querySelector(".chart");

    const ChartLib = await this.getLib();
    this._chart = new ChartLib(ctx, {
      type: "bar",
      data: this.parseData(),
      options: {
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
  }

  // Custom element removed from page
  disconnectedCallback() {
    // ...
  }

  // Custom element moved to new page
  addoptedCallback() {
    // ...
  }

  // attribute change callback
  attributeChangedCallback(name, oldValue, newValue) {
    console.log(`Attribute ${name} has changed.`);
  }
}
