export interface ISpeed {
  year: number;
  location: string;
  weekdayKmh: number;
  weekdayMs: number;
  saturdayKmh: number;
  saturdayMs: number;
  sundayKmh: number;
  sundayMs: number;
}

export interface IBasicLink {
  routeId: string;
  startStationId: string;
  endStationId: string;
  tripTime: number;
  tripDistance: number;
  stationOrder: number;
}
