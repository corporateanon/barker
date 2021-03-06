/* Do not change, this code is generated from Golang structs */


export enum DeliveryState {
    progress = 1,
    fail = 3,
    success = 2,
}
export interface Bot {
    ID?: number;
    Title?: string;
    Token?: string;
    RRAccessTime?: string;
    RRPossiblyEmpty?: boolean;
}
export interface Campaign {
    ID?: number;
    BotID?: number;
    Title?: string;
    Message?: string;
    Active?: boolean;
}
export interface CampaignAggregatedStatistics {
    Users?: number;
    Delivered?: number;
    Errors?: number;
    Pending?: number;
    TimedOut?: number;
}
export interface User {
    FirstName?: string;
    LastName?: string;
    DisplayName?: string;
    UserName?: string;
    TelegramID?: number;
    BotID?: number;
}
export interface Delivery {
    CampaignID?: number;
    BotID?: number;
    TelegramID?: number;
    State?: DeliveryState;
}
export interface PaginatorRequest {
    Page?: number;
    Size?: number;
}
export interface PaginatorResponse {
    Page?: number;
    Size?: number;
    Total?: number;
    TotalItems?: number;
}
export interface DeliveryTakeResult {
    Delivery?: Delivery;
    Campaign?: Campaign;
    User?: User;
}