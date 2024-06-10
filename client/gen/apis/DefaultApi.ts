/* tslint:disable */
/* eslint-disable */
/**
 * Roast Advisor API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  Review,
} from '../models/index';
import {
    ReviewFromJSON,
    ReviewToJSON,
} from '../models/index';

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Responds with the list of all reviews as JSON.
     * Get array of reviews
     */
    async reviewsGetRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<Review>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/reviews`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(ReviewFromJSON));
    }

    /**
     * Responds with the list of all reviews as JSON.
     * Get array of reviews
     */
    async reviewsGet(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<Review>> {
        const response = await this.reviewsGetRaw(initOverrides);
        return await response.value();
    }

}
