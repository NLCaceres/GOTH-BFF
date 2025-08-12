export const elementIs = (element, name) => element.tagName == name;

export const hasClass = (element, className) =>
  element.classList.contains(className);

export const elementWithClass = (element, name, className) =>
  elementIs(element, name) && hasClass(element, className);
