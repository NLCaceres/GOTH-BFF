/** Takes a HTML element and matches the `name` arg with the element's `tagName`
 * @param {string} name - The expected HTML `tagName`, always in all upper-case form
 * @return {boolean} `true` if the `tagName` matches the input `name` */
export const elementIs = (element, name) =>
  element.tagName == name.toUpperCase();

/** Takes a HTML element and checks its CSS `classList` for a matching `className`
 * @param {string} className - CSS class name
 * @return {boolean} `true` if the class name is found in the element `classList` */
export const hasClass = (element, className) =>
  element.classList.contains(className);

/**
 * Takes a HTML element and matches the `name` arg with the element's `tagName`
 * AND also checks that the element has the input CSS `class`
 * @param {string} name - The expected HTML `tagName`, always in all upper-case form
 * @param {string} className - CSS class name
 * @return {boolean} `true` if both the `name` and CSS `className` match the property
 * values found in the HTML element `tagName` and `classList`, respectively */
export const elementWithClass = (element, name, className) =>
  elementIs(element, name) && hasClass(element, className);
