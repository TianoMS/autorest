/**
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for
 * license information.
 *
 * Code generated by Microsoft (R) AutoRest Code Generator.
 * Changes may cause incorrect behavior and will be lost if the code is
 * regenerated.
 */

package fixtures.bodydatetime;

import com.microsoft.rest.ServiceCall;
import com.microsoft.rest.ServiceCallback;
import com.microsoft.rest.ServiceResponse;
import fixtures.bodydatetime.models.ErrorException;
import java.io.IOException;
import org.joda.time.DateTime;

/**
 * An instance of this class provides access to all the operations defined
 * in Datetimes.
 */
public interface Datetimes {
    /**
     * Get null datetime value.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getNull() throws ErrorException, IOException;

    /**
     * Get null datetime value.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getNullAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Get invalid datetime value.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getInvalid() throws ErrorException, IOException;

    /**
     * Get invalid datetime value.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getInvalidAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Get overflow datetime value.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getOverflow() throws ErrorException, IOException;

    /**
     * Get overflow datetime value.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getOverflowAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Get underflow datetime value.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getUnderflow() throws ErrorException, IOException;

    /**
     * Get underflow datetime value.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getUnderflowAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Put max datetime value 9999-12-31T23:59:59.9999999Z.
     *
     * @param datetimeBody the DateTime value
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @throws IllegalArgumentException exception thrown from invalid parameters
     * @return the {@link ServiceResponse} object if successful.
     */
    ServiceResponse<Void> putUtcMaxDateTime(DateTime datetimeBody) throws ErrorException, IOException, IllegalArgumentException;

    /**
     * Put max datetime value 9999-12-31T23:59:59.9999999Z.
     *
     * @param datetimeBody the DateTime value
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<Void> putUtcMaxDateTimeAsync(DateTime datetimeBody, final ServiceCallback<Void> serviceCallback) throws IllegalArgumentException;

    /**
     * Get max datetime value 9999-12-31t23:59:59.9999999z.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getUtcLowercaseMaxDateTime() throws ErrorException, IOException;

    /**
     * Get max datetime value 9999-12-31t23:59:59.9999999z.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getUtcLowercaseMaxDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Get max datetime value 9999-12-31T23:59:59.9999999Z.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getUtcUppercaseMaxDateTime() throws ErrorException, IOException;

    /**
     * Get max datetime value 9999-12-31T23:59:59.9999999Z.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getUtcUppercaseMaxDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Put max datetime value with positive numoffset 9999-12-31t23:59:59.9999999+14:00.
     *
     * @param datetimeBody the DateTime value
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @throws IllegalArgumentException exception thrown from invalid parameters
     * @return the {@link ServiceResponse} object if successful.
     */
    ServiceResponse<Void> putLocalPositiveOffsetMaxDateTime(DateTime datetimeBody) throws ErrorException, IOException, IllegalArgumentException;

    /**
     * Put max datetime value with positive numoffset 9999-12-31t23:59:59.9999999+14:00.
     *
     * @param datetimeBody the DateTime value
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<Void> putLocalPositiveOffsetMaxDateTimeAsync(DateTime datetimeBody, final ServiceCallback<Void> serviceCallback) throws IllegalArgumentException;

    /**
     * Get max datetime value with positive num offset 9999-12-31t23:59:59.9999999+14:00.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getLocalPositiveOffsetLowercaseMaxDateTime() throws ErrorException, IOException;

    /**
     * Get max datetime value with positive num offset 9999-12-31t23:59:59.9999999+14:00.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getLocalPositiveOffsetLowercaseMaxDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Get max datetime value with positive num offset 9999-12-31T23:59:59.9999999+14:00.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getLocalPositiveOffsetUppercaseMaxDateTime() throws ErrorException, IOException;

    /**
     * Get max datetime value with positive num offset 9999-12-31T23:59:59.9999999+14:00.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getLocalPositiveOffsetUppercaseMaxDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Put max datetime value with positive numoffset 9999-12-31t23:59:59.9999999-14:00.
     *
     * @param datetimeBody the DateTime value
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @throws IllegalArgumentException exception thrown from invalid parameters
     * @return the {@link ServiceResponse} object if successful.
     */
    ServiceResponse<Void> putLocalNegativeOffsetMaxDateTime(DateTime datetimeBody) throws ErrorException, IOException, IllegalArgumentException;

    /**
     * Put max datetime value with positive numoffset 9999-12-31t23:59:59.9999999-14:00.
     *
     * @param datetimeBody the DateTime value
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<Void> putLocalNegativeOffsetMaxDateTimeAsync(DateTime datetimeBody, final ServiceCallback<Void> serviceCallback) throws IllegalArgumentException;

    /**
     * Get max datetime value with positive num offset 9999-12-31T23:59:59.9999999-14:00.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getLocalNegativeOffsetUppercaseMaxDateTime() throws ErrorException, IOException;

    /**
     * Get max datetime value with positive num offset 9999-12-31T23:59:59.9999999-14:00.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getLocalNegativeOffsetUppercaseMaxDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Get max datetime value with positive num offset 9999-12-31t23:59:59.9999999-14:00.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getLocalNegativeOffsetLowercaseMaxDateTime() throws ErrorException, IOException;

    /**
     * Get max datetime value with positive num offset 9999-12-31t23:59:59.9999999-14:00.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getLocalNegativeOffsetLowercaseMaxDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Put min datetime value 0001-01-01T00:00:00Z.
     *
     * @param datetimeBody the DateTime value
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @throws IllegalArgumentException exception thrown from invalid parameters
     * @return the {@link ServiceResponse} object if successful.
     */
    ServiceResponse<Void> putUtcMinDateTime(DateTime datetimeBody) throws ErrorException, IOException, IllegalArgumentException;

    /**
     * Put min datetime value 0001-01-01T00:00:00Z.
     *
     * @param datetimeBody the DateTime value
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<Void> putUtcMinDateTimeAsync(DateTime datetimeBody, final ServiceCallback<Void> serviceCallback) throws IllegalArgumentException;

    /**
     * Get min datetime value 0001-01-01T00:00:00Z.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getUtcMinDateTime() throws ErrorException, IOException;

    /**
     * Get min datetime value 0001-01-01T00:00:00Z.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getUtcMinDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Put min datetime value 0001-01-01T00:00:00+14:00.
     *
     * @param datetimeBody the DateTime value
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @throws IllegalArgumentException exception thrown from invalid parameters
     * @return the {@link ServiceResponse} object if successful.
     */
    ServiceResponse<Void> putLocalPositiveOffsetMinDateTime(DateTime datetimeBody) throws ErrorException, IOException, IllegalArgumentException;

    /**
     * Put min datetime value 0001-01-01T00:00:00+14:00.
     *
     * @param datetimeBody the DateTime value
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<Void> putLocalPositiveOffsetMinDateTimeAsync(DateTime datetimeBody, final ServiceCallback<Void> serviceCallback) throws IllegalArgumentException;

    /**
     * Get min datetime value 0001-01-01T00:00:00+14:00.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getLocalPositiveOffsetMinDateTime() throws ErrorException, IOException;

    /**
     * Get min datetime value 0001-01-01T00:00:00+14:00.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getLocalPositiveOffsetMinDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

    /**
     * Put min datetime value 0001-01-01T00:00:00-14:00.
     *
     * @param datetimeBody the DateTime value
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @throws IllegalArgumentException exception thrown from invalid parameters
     * @return the {@link ServiceResponse} object if successful.
     */
    ServiceResponse<Void> putLocalNegativeOffsetMinDateTime(DateTime datetimeBody) throws ErrorException, IOException, IllegalArgumentException;

    /**
     * Put min datetime value 0001-01-01T00:00:00-14:00.
     *
     * @param datetimeBody the DateTime value
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<Void> putLocalNegativeOffsetMinDateTimeAsync(DateTime datetimeBody, final ServiceCallback<Void> serviceCallback) throws IllegalArgumentException;

    /**
     * Get min datetime value 0001-01-01T00:00:00-14:00.
     *
     * @throws ErrorException exception thrown from REST call
     * @throws IOException exception thrown from serialization/deserialization
     * @return the DateTime object wrapped in {@link ServiceResponse} if successful.
     */
    ServiceResponse<DateTime> getLocalNegativeOffsetMinDateTime() throws ErrorException, IOException;

    /**
     * Get min datetime value 0001-01-01T00:00:00-14:00.
     *
     * @param serviceCallback the async ServiceCallback to handle successful and failed responses.
     * @throws IllegalArgumentException thrown if callback is null
     * @return the {@link ServiceCall} object
     */
    ServiceCall<DateTime> getLocalNegativeOffsetMinDateTimeAsync(final ServiceCallback<DateTime> serviceCallback) throws IllegalArgumentException;

}
