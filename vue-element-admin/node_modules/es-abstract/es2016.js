'use strict';

var hasSymbols = require('has-symbols')();

var ES2015 = require('./es2015');
var assign = require('./helpers/assign');

var callBound = require('./helpers/callBound');

var $arrayPush = callBound('Array.prototype.push');
var $arraySlice = callBound('Array.prototype.slice');
var $arrayJoin = callBound('Array.prototype.join');

var ES2016 = assign(assign({}, ES2015), {
	// https://www.ecma-international.org/ecma-262/7.0/#sec-samevaluenonnumber
	SameValueNonNumber: function SameValueNonNumber(x, y) {
		if (typeof x === 'number' || typeof x !== typeof y) {
			throw new TypeError('SameValueNonNumber requires two non-number values of the same type.');
		}
		return this.SameValue(x, y);
	},

	// https://www.ecma-international.org/ecma-262/7.0/#sec-iterabletoarraylike
	IterableToArrayLike: function IterableToArrayLike(items) {
		var usingIterator;
		if (hasSymbols) {
			usingIterator = this.GetMethod(items, Symbol.iterator);
		} else if (this.IsArray(items)) {
			usingIterator = function () {
				var i = -1;
				var arr = this; // eslint-disable-line no-invalid-this
				return {
					next: function () {
						i += 1;
						return {
							done: i >= arr.length,
							value: arr[i]
						};
					}
				};
			};
		} else if (this.Type(items) === 'String') {
			var ES = this;
			usingIterator = function () {
				var i = 0;
				return {
					next: function () {
						var nextIndex = ES.AdvanceStringIndex(items, i, true);
						var value = $arrayJoin($arraySlice(items, i, nextIndex), '');
						i = nextIndex;
						return {
							done: nextIndex > items.length,
							value: value
						};
					}
				};
			};
		}
		if (typeof usingIterator !== 'undefined') {
			var iterator = this.GetIterator(items, usingIterator);
			var values = [];
			var next = true;
			while (next) {
				next = this.IteratorStep(iterator);
				if (next) {
					var nextValue = this.IteratorValue(next);
					$arrayPush(values, nextValue);
				}
			}
			return values;
		}

		return this.ToObject(items);
	}
});

module.exports = ES2016;
