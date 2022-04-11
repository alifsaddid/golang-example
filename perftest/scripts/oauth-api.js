
/* eslint-disable no-undef */
/* global, __ENV */

import { check, fail } from 'k6'
import http from 'k6/http'

export const getAll = (scenarioOptions) => {
    const resp = http.get('http://host-1906398250-port-58250.proxy.infralabs.cs.ui.ac.id/roles')

    scenarioOptions.trends.searchTrend.add(resp.timings.duration)

    check(resp, {
        'status is 200': resp => resp.status == 200
    })

    if (resp.status !== 200) {
        fail(`response is ${resp.status} - ${resp.body}`)
    }
}

export const create = (params, scenarioOptions) => {
    const reqBody = {
        name: params.name,
        permissions: params.permissions
    }

    const resp = http.post(
        'http://host-1906398250-port-58250.proxy.infralabs.cs.ui.ac.id/roles',
        JSON.stringify(reqBody),
        params
    )

    scenarioOptions.trends.searchTrend.add(resp.timings.duration)

    check(resp, {
        'status is 200': resp => resp.status == 200
    })

    if (resp.status !== 200) {
        fail(`response is ${resp.status} - ${resp.body} - ${JSON.stringify(reqBody)}`)
    }
}

export const update = (params, scenarioOptions) => {
    const reqBody = {
        name: params.name,
        permissions: params.permissions
    }

    const resp = http.put(
        'http://host-1906398250-port-58250.proxy.infralabs.cs.ui.ac.id/roles/567',
        JSON.stringify(reqBody),
        params
    )

    scenarioOptions.trends.searchTrend.add(resp.timings.duration)

    check(resp, {
        'status is 200': resp => resp.status == 200
    })

    if (resp.status !== 200) {
        fail(`response is ${resp.status} - ${resp.body} - ${JSON.stringify(reqBody)}`)
    }
}

export const test = (params, scenarioOptions) => {
    if (params.operation == 10000) {
        create(params, scenarioOptions)
    } else if (params.operation < 10000 && params.operation >= 9900) {
        update(params, scenarioOptions)
    } else {
        getAll(scenarioOptions)
    }
}