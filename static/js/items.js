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

function itemLinkOverflowing() {
  for (const item of document.getElementsByClassName("itemLink")) {
    if ((item.scrollHeight - item.offsetHeight) > 5) {
      item.lastElementChild.classList.add("overflowing");
    }
  }
}

// Runs ONLY upon resize (so could possibly NOT run)
let resizeTimeout;
window.addEventListener("resize", function() {
  window.clearTimeout(resizeTimeout);
  resizeTimeout = window.setTimeout(itemLinkOverflowing, 500);
});

// Runs as soon the HTML is parsed BUT often before it's fully constructed
window.addEventListener("DOMContentLoaded", function() {
  itemLinkOverflowing()
  for (const checkBox of document.getElementsByClassName("checkbox-button")) {
    checkBox.addEventListener('click', checkBoxClicked)
  }
});
