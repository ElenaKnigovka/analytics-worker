// types.ts
import { Maybe } from './maybe';

export interface IIdentifier {
  readonly id: string;
}

export interface IEvent {
  readonly id: string;
  readonly name: string;
  readonly timestamp: Date;
  readonly attributes: { [key: string]: string };
}

export interface IEventAttributes {
  [key: string]: string | Maybe<string>;
}

export interface ICategoryAttributes {
  [key: string]: string | Maybe<string>;
}

export type EventAttributes = Record<string, string | Maybe<string>>;
export type CategoryAttributes = Record<string, string | Maybe<string>>;

export interface IEventCategory {
  readonly id: string;
  readonly name: string;
  readonly attributes: CategoryAttributes;
}

export type EventData = Record<string, string>;

export interface IAnalyticsData {
  readonly events: IEvent[];
  readonly categories: IEventCategory[];
}

export class AnalyticsData implements IAnalyticsData {
  public readonly events: IEvent[];
  public readonly categories: IEventCategory[];

  constructor(events: IEvent[], categories: IEventCategory[]) {
    this.events = events;
    this.categories = categories;
  }
}