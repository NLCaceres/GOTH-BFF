import { debounce } from "./helpers.js";

/** Checks all `.itemLink` elems, adding CSS marker class if <p> inside overflows */
function itemLinkOverflowing() {
  for (const item of document.getElementsByClassName("itemLink")) {
    if (item.scrollHeight - item.offsetHeight > 5) {
      item.lastElementChild.classList.add("overflowing");
    } else {
      item.lastElementChild.classList.remove("overflowing");
    }
  }
}

// Runs ONLY upon resize (so could possibly NEVER run)
window.addEventListener("resize", debounce(itemLinkOverflowing, 500));

/** Syncs the ID container color with the checkbox background color in `.listItem` elem */
function syncItemColoring(e) {
  const checkbox = e.currentTarget;
  const id =
    checkbox.parentElement?.previousElementSibling?.previousElementSibling
      ?.firstElementChild;
  const checkboxBackground = checkbox.previousElementSibling?.firstElementChild;
  if (!checkbox.classList.contains("check")) {
    id?.classList.replace("bg-green", "bg-red");
    checkboxBackground?.setAttribute("stroke", "red");
  } else {
    id?.classList.replace("bg-red", "bg-green");
    checkboxBackground?.setAttribute("stroke", "green");
  }
}

// Runs as soon the HTML is parsed BUT often before it's fully constructed
window.addEventListener("DOMContentLoaded", function () {
  itemLinkOverflowing();
  for (const checkBox of document.getElementsByClassName("checkbox-button")) {
    checkBox.addEventListener("click", syncItemColoring);
  }
});
