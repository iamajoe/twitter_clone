export class Test extends HTMLElement {
  // connect component
  connectedCallback() {
    console.log("connectedCallback");
    this.textContent = "Hello World!";
  }
}
