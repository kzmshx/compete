/**
 * @param {number} n
 * @returns {HTMLElement}
 */
const findNthInputElement = (n) =>
  Array.from(document.getElementsByTagName("h3"))
    .find((el) => el.innerText.includes(`入力例 ${n}`))
    ?.parentElement?.getElementsByTagName("pre")[0];

/**
 * @param {number} n
 * @returns {HTMLElement}
 */
const findNthOutputElement = (n) =>
  Array.from(document.getElementsByTagName("h3"))
    .find((el) => el.innerText.includes(`出力例 ${n}`))
    ?.parentElement?.getElementsByTagName("pre")[0];

/**
 * @param {number} ms
 * @returns {Promise<void>}
 */
const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

/**
 * @param {string} filename
 * @param {string} content
 * @returns {Promise<void>}
 */
const downloadFile = async (filename, content) => {
  const blob = new Blob([content], { type: "text/plain" });
  const url = URL.createObjectURL(blob);
  const file = document.createElement("a");
  file.href = url;
  file.download = filename;
  file.click();
  URL.revokeObjectURL(url);
  await sleep(200);
};

/**
 * Download AtCoder Problem Test Cases
 */
const main = async () => {
  for (let i = 1; ; i++) {
    const [inputElem, outputElem] = [
      findNthInputElement(i),
      findNthOutputElement(i),
    ];
    if (!inputElem || !outputElem) {
      break;
    }
    await downloadFile(`testcase_${i}_input.txt`, inputElem.innerText);
    await downloadFile(`testcase_${i}_output.txt`, outputElem.innerText);
  }
};

await main();
