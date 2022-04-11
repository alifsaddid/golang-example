/* global open, __ENV */

import { Trend } from 'k6/metrics'
import { group } from 'k6'
import {
    randomIntBetween,
    uuidv4,
    randomString
} from 'https://jslib.k6.io/k6-utils/1.0.0/index.js'
import { test } from './oauth-api.js'

let stages
if (`${__ENV.VUS}` === undefined || `${__ENV.VUS}` === 'undefined') {
    stages = [{ duration: '24h', target: 50 }]
} else {
    stages = [
        { duration: '2.5m', target: `${__ENV.VUS}` },
        { duration: `${(__ENV.DURATION) - 5}m`, target: `${__ENV.VUS}` },
        { duration: '2.5m', target: 0 }
    ]
}

export const options = {
    stages: stages
}

const scenario = {
    config: {
        baseURL: `http://${__ENV.API_HOSTNAME}/`
    },
    trends: {
        searchTrend: new Trend('/oauth/roles')
    }
}

export default function () {
    // default header parameter
    const params = {
        tags: {
            name: '',
            host: `${__ENV.API_HOSTNAME}`
        },
        name: randomString(5),
        permissions: [randomIntBetween(1, 3), randomIntBetween(4, 6), randomIntBetween(7, 10)],
        operation: randomIntBetween(1, 10000)   
    }

    // trigger the test by groups sequentialy
    group('getAll', () => {
        test(params, scenario)
    })
}