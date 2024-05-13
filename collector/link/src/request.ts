export const request = async <T>(
  url: (start: number, end: number) => string,
  start: number,
  end: number,
) => {
  const resp = await fetch(url(start, end), {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return (await resp.json()) as T;
};
