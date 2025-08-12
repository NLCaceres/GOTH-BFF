/* Helper functions */
const elementIs = (element, name) => element.tagName == name;
const hasClass = (element, className) => element.classList.contains(className);
const elementWithClass = (element, name, className) =>
  elementIs(element, name) && hasClass(element, className);

/* Mouse event handlers to simulate a CSS hover effect */
function mouseEntered(event) {
  if (!elementWithClass(event.target, "DIV", "itemLink")) {
    return;
  }
  event.target.dataset.expanding = true;
  delete event.target.dataset.closed;
}
function mouseLeft(event) {
  if (!elementWithClass(event.target, "DIV", "itemLink")) {
    return;
  }
  event.target.dataset.closed = true;
  delete event.target.dataset.expanding;
}

/* Transition Event handlers to help smoothen the CSS hover effect */
function transitionStarted(event) {
  if (!elementWithClass(event.target, "LI", "listItem")) {
    return;
  }
  const link = event.target.querySelector(".itemLink");
  if (link.dataset.expanding === "true") {
    link.classList.add("no-ellipsis");
  }
}
function transitionEnded(event) {
  if (!elementWithClass(event.target, "LI", "listItem")) {
    return;
  }
  const link = event.target.querySelector(".itemLink");
  if (link.dataset.closed === "true") {
    link.classList.remove("no-ellipsis");
  }
}

/* Runs when <script> loads this file, adding event handlers IF it finds matching elems */
for (const item of document.getElementsByClassName("items")) {
  item.addEventListener("mouseenter", mouseEntered, true);
  item.addEventListener("mouseleave", mouseLeft, true);
  item.addEventListener("transitionstart", transitionStarted, true);
  item.addEventListener("transitionend", transitionEnded, true);
}
