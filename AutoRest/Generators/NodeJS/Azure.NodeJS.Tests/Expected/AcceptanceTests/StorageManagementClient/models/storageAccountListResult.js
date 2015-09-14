/*
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for
 * license information.
 * 
 * Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
 * Changes may cause incorrect behavior and will be lost if the code is
 * regenerated.
 */

'use strict';

var models = require('./index');

var util = require('util');

/**
 * @class
 * Initializes a new instance of the StorageAccountListResult class.
 * @constructor
 * The list storage accounts operation response.
 * @member {array} [value] Gets the list of storage accounts and their
 * properties.
 * 
 * @member {string} [nextLink] Gets the link to the next set of results.
 * Currently this will always be empty as the API does not support pagination.
 * 
 */
function StorageAccountListResult(parameters) {
  if (parameters !== null && parameters !== undefined) {
    if (parameters.value !== null && parameters.value !== undefined) {
      var initializedParametersvalue = [];
      parameters.value.forEach(function(element) {
        if (element !== null && element !== undefined) {
          element = new models['StorageAccount'](element);
        }
        initializedParametersvalue.push(element);
      });
      this.value = initializedParametersvalue;
    }
    if (parameters.nextLink !== null && parameters.nextLink !== undefined) {
      this.nextLink = parameters.nextLink;
    }
  }    
}


/**
 * Validate the payload against the StorageAccountListResult schema
 *
 * @param {JSON} payload
 *
 */
StorageAccountListResult.prototype.serialize = function () {
  var payload = {};
  if (util.isArray(this['value'])) {
    for (var i = 0; i < this['value'].length; i++) {
      if (this['value'][i]) {
        payload['value'][i] = this['value'][i].serialize();
      }
    }
  }

  if (this['nextLink'] !== null && this['nextLink'] !== undefined) {
    if (typeof this['nextLink'].valueOf() !== 'string') {
      throw new Error('this[\'nextLink\'] must be of type string.');
    }
    payload['nextLink'] = this['nextLink'];
  }
};

/**
 * Deserialize the instance to StorageAccountListResult schema
 *
 * @param {JSON} instance
 *
 */
StorageAccountListResult.prototype.deserialize = function (instance) {
  if (instance) {
    if (instance['value'] !== null && instance['value'] !== undefined) {
      var deserializedInstancevalue = [];
      instance['value'].forEach(function(element1) {
        if (element1 !== null && element1 !== undefined) {
          element1 = new models['StorageAccount']().deserialize(element1);
        }
        deserializedInstancevalue.push(element1);
      });
      this['value'] = deserializedInstancevalue;
    }

    if (instance['nextLink'] !== null && instance['nextLink'] !== undefined) {
      this['nextLink'] = instance['nextLink'];
    }
  }
};

module.exports = StorageAccountListResult;
