import { addSeconds } from 'date-fns'

export const distanceDateFromSeconds = (date, nextSeconds) => {
  return formatDate(
    addSeconds(
      new Date(date), nextSeconds,
    )
  );
};

export const formatDate = (date) => {
  const d = new Date(date);
  return new Intl.DateTimeFormat('en-GB',{
    year: 'numeric', month: 'numeric', day: 'numeric',
    hour: 'numeric', minute: 'numeric', second: 'numeric',
    hour12: false,
    timeZone: 'Asia/Bangkok',
  }).format(d);
};
