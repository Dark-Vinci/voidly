export function fi<T = any>(cond: boolean, value: T): T | null {
  if (!cond) {
    return null;
  }

  return value;
}
