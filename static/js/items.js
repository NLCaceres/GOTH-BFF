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

/* Following code runs when the `<script>` is loaded */
let resizeTimeout;
window.addEventListener("resize", function() {
  window.clearTimeout(resizeTimeout);
  resizeTimeout = window.setTimeout(function() {
    for (const item of document.getElementsByClassName("itemLink")) {
      if ((item.scrollHeight - item.offsetHeight) > 5) {
        item.lastElementChild.classList.add("overflowing");
      }
    }
  }, 500);
})

for (const checkBox of document.getElementsByClassName("checkbox-button")) {
  checkBox.addEventListener('click', checkBoxClicked)
}
