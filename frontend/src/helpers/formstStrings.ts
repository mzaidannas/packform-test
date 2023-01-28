export function snakeToPascal(str: string) {
  return str
    .split('_')
    .map((substr: string) => substr.charAt(0).toUpperCase() + substr.slice(1))
    .join('');
}
