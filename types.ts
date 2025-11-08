// types.ts
export type AnalyticsEvent = {
  id: string;
  timestamp: number;
  type: 'page_view' | 'click' | 'scroll';
  data: {
    url: string;
    userAgent: string;
    screenWidth: number;
    screenHeight: number;
  };
};

export type AnalyticsEventData = {
  events: AnalyticsEvent[];
  lastSync: number | null;
};

export type WorkerMessage = {
  type: 'sync' | 'event';
  data: AnalyticsEventData | AnalyticsEvent;
};

export enum WorkerMessageType {
  Sync = 'sync',
  Event = 'event',
}

export interface WorkerMessageHandler {
  (message: WorkerMessage): void;
}
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WorkerErrorHandler = (error: any) => void;