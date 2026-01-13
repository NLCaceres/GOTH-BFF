import { elementWithClass } from "./helpers.js";

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

function checkBoxClicked(e) {
  const item = e.currentTarget;
  const id = item.parentElement?.previousElementSibling?.previousElementSibling?.firstElementChild;
  const check = item.previousElementSibling?.firstElementChild;
  if (item.classList.contains('check')) {
    id?.classList.replace('bg-green', 'bg-red');
    check?.setAttribute('stroke', 'red');
  }
  else {
    id?.classList.replace('bg-red', 'bg-green');
    check?.setAttribute('stroke', 'green');
  }
}

/* Runs when <script> loads this file, adding event handlers IF it finds matching elems */
for (const item of document.getElementsByClassName("items")) {
  item.addEventListener("mouseenter", mouseEntered, true);
  item.addEventListener("mouseleave", mouseLeft, true);
  item.addEventListener("transitionstart", transitionStarted, true);
  item.addEventListener("transitionend", transitionEnded, true);
}

for (const checkBox of document.getElementsByClassName("checkbox-button")) {
  checkBox.addEventListener('click', checkBoxClicked)
}
