import React, { useState, useMemo } from 'react';
import { Dialog } from 'primereact/dialog';
import { Button } from 'primereact/button';
import { ProgressBar } from 'primereact/progressbar';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { Dropdown } from 'primereact/dropdown';
import { Calendar } from 'primereact/calendar';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import { typography, colors, textStyles } from '../styles/typography';

interface CreateLoadDialogProps {
    visible: boolean;
    onHide: () => void;
    onCreateLoad: (loadData: any) => Promise<any>;
    isLoading: boolean;
}

// Shared styles - using functions to avoid mutation issues
const getInputTextStyle = () => ({
    padding: '0.75rem 1rem',
    border: `1px solid ${colors.border.light}`,
    borderRadius: '8px',
    fontSize: typography.sizes.base,
    fontFamily: typography.fonts.primary,
    backgroundColor: colors.background.primary,
    color: colors.text.primary,
    transition: 'all 0.2s ease',
    width: '100%'
});

const getLabelStyle = () => ({
    ...textStyles.label,
    display: 'block',
    marginBottom: '0.5rem'
});

const getErrorStyle = () => ({
    color: colors.status.error, 
    fontSize: typography.sizes.xs, 
    fontWeight: typography.weights.medium,
    display: 'block',
    marginTop: '4px',
    fontFamily: typography.fonts.primary
});

const getGridContainerStyle = () => ({
    background: colors.background.primary,
    display: 'grid',
    gridTemplateColumns: '1fr 1fr',
    gap: '1.5rem',
    marginBottom: '1rem'
});

const CreateLoadDialog: React.FC<CreateLoadDialogProps> = ({ visible, onHide, onCreateLoad, isLoading }) => {
    const [currentStep, setCurrentStep] = useState(1);

    // Validation schema and options
    const validationSchema = useMemo(() => Yup.object({
        // Step 1 fields
        pickupName: Yup.string().required('Pickup name is required'),
        pickupCity: Yup.string().required('Pickup city is required'),
        pickupState: Yup.string().required('Pickup state is required'),
        consigneeName: Yup.string().required('Consignee name is required'),
        consigneeCity: Yup.string().required('Consignee city is required'),
        consigneeState: Yup.string().required('Consignee state is required'),
        // Step 2 fields
        pickupDate: Yup.date().required('Pickup date is required'),
        deliveryDate: Yup.date().required('Delivery date is required'),
        status: Yup.string().required('Status is required'),
        customerName: Yup.string().required('Customer name is required'),
        customerTMSId: Yup.string().required('Customer TMS ID is required'),
        totalWeight: Yup.number().required('Weight is required').min(1, 'Weight must be greater than 0'),
        minTemp: Yup.number().required('Minimum temperature is required'),
        maxTemp: Yup.number().required('Maximum temperature is required'),
    }), []);

    const statusOptions = useMemo(() => [
        { label: 'Tendered', value: 'Tendered' },
        { label: 'Covered', value: 'Covered' },
        { label: 'In Transit', value: 'In Transit' },
        { label: 'Delivered', value: 'Delivered' }
    ], []);

    const formik = useFormik({
        initialValues: {
            pickupName: '', pickupCity: '', pickupState: '',
            consigneeName: '', consigneeCity: '', consigneeState: '',
            pickupDate: null as Date | null, deliveryDate: null as Date | null,
            status: '', customerName: '', customerTMSId: '',
            totalWeight: 0, minTemp: 32, maxTemp: 75,
        },
        validationSchema,
        onSubmit: async (values) => {
            const loadData = {
                pickup: {
                    name: values.pickupName,
                    apptTime: values.pickupDate?.toISOString() || '',
                    city: values.pickupCity,
                    state: values.pickupState,
                    country: 'USA'
                },
                consignee: {
                    name: values.consigneeName,
                    apptTime: values.deliveryDate?.toISOString() || '',
                    city: values.consigneeCity,
                    state: values.consigneeState,
                    country: 'USA'
                },
                status: values.status,
                customer: {
                    name: values.customerName,
                    externalTMSId: values.customerTMSId
                },
                specifications: {
                    minTempFahrenheit: values.minTemp,
                    maxTempFahrenheit: values.maxTemp
                },
                totalWeight: values.totalWeight
            };

            await onCreateLoad(loadData);
        },
    });

    const handleNext = async () => {
        const step1Fields = ['pickupName', 'pickupCity', 'pickupState', 'consigneeName', 'consigneeCity', 'consigneeState'];
        const step2Fields = ['pickupDate', 'deliveryDate', 'status', 'customerName', 'customerTMSId', 'totalWeight', 'minTemp', 'maxTemp'];
        const fieldsToValidate = currentStep === 1 ? step1Fields : step2Fields;
        
        const errors = await formik.validateForm();
        const hasErrors = fieldsToValidate.some(field => errors[field as keyof typeof errors]);
        
        if (!hasErrors) {
            currentStep === 1 ? setCurrentStep(2) : formik.handleSubmit();
        } else {
            const touched = fieldsToValidate.reduce((acc, field) => ({ ...acc, [field]: true }), {});
            formik.setTouched(touched);
        }
    };

    const handlePrevious = () => {
        setCurrentStep(1);
    };

    const handleDialogClose = () => {
        onHide();
        setCurrentStep(1);
        formik.resetForm();
    };

    const renderStep1 = () => (
        <div style={getGridContainerStyle()}>
            <div>
                <label htmlFor="pickupName" style={getLabelStyle()}>
                    Pickup Name
                </label>
                <InputText
                    id="pickupName"
                    name="pickupName"
                    value={formik.values.pickupName}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.pickupName && formik.errors.pickupName ? 'p-invalid' : ''}`}
                    placeholder="Enter pickup name"
                    style={getInputTextStyle()}
                />
                {formik.touched.pickupName && formik.errors.pickupName && (
                    <small style={getErrorStyle()}>
                        {formik.errors.pickupName}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="consigneeName" style={getLabelStyle()}>
                    Consignee Name
                </label>
                <InputText
                    id="consigneeName"
                    name="consigneeName"
                    value={formik.values.consigneeName}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.consigneeName && formik.errors.consigneeName ? 'p-invalid' : ''}`}
                    placeholder="Enter consignee name"
                    style={getInputTextStyle()}
                />
                {formik.touched.consigneeName && formik.errors.consigneeName && (
                    <small style={getErrorStyle()}>
                        {formik.errors.consigneeName}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="pickupCity" style={getLabelStyle()}>
                    Pickup City
                </label>
                <InputText
                    id="pickupCity"
                    name="pickupCity"
                    value={formik.values.pickupCity}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.pickupCity && formik.errors.pickupCity ? 'p-invalid' : ''}`}
                    placeholder="Enter pickup city"
                    style={getInputTextStyle()}
                />
                {formik.touched.pickupCity && formik.errors.pickupCity && (
                    <small style={getErrorStyle()}>
                        {formik.errors.pickupCity}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="consigneeCity" style={getLabelStyle()}>
                    Consignee City
                </label>
                <InputText
                    id="consigneeCity"
                    name="consigneeCity"
                    value={formik.values.consigneeCity}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.consigneeCity && formik.errors.consigneeCity ? 'p-invalid' : ''}`}
                    placeholder="Enter consignee city"
                    style={getInputTextStyle()}
                />
                {formik.touched.consigneeCity && formik.errors.consigneeCity && (
                    <small style={getErrorStyle()}>
                        {formik.errors.consigneeCity}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="pickupState" style={getLabelStyle()}>
                    Pickup State
                </label>
                <InputText
                    id="pickupState"
                    name="pickupState"
                    value={formik.values.pickupState}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.pickupState && formik.errors.pickupState ? 'p-invalid' : ''}`}
                    placeholder="Enter pickup state"
                    style={getInputTextStyle()}
                />
                {formik.touched.pickupState && formik.errors.pickupState && (
                    <small style={getErrorStyle()}>
                        {formik.errors.pickupState}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="consigneeState" style={getLabelStyle()}>
                    Consignee State
                </label>
                <InputText
                    id="consigneeState"
                    name="consigneeState"
                    value={formik.values.consigneeState}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.consigneeState && formik.errors.consigneeState ? 'p-invalid' : ''}`}
                    placeholder="Enter consignee state"
                    style={getInputTextStyle()}
                />
                {formik.touched.consigneeState && formik.errors.consigneeState && (
                    <small style={getErrorStyle()}>
                        {formik.errors.consigneeState}
                    </small>
                )}
            </div>
        </div>
    );

    const renderStep2 = () => (
        <div style={getGridContainerStyle()}>
            <div>
                <label htmlFor="pickupDate" style={getLabelStyle()}>
                    Pickup Date
                </label>
                <Calendar
                    id="pickupDate"
                    name="pickupDate"
                    value={formik.values.pickupDate}
                    onChange={(e) => formik.setFieldValue('pickupDate', e.value)}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.pickupDate && formik.errors.pickupDate ? 'p-invalid' : ''}`}
                    placeholder="Select pickup date"
                    showIcon
                    showTime
                    style={{
                        backgroundColor: colors.background.primary,
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        width: '100%'
                    }}
                    inputStyle={{
                        padding: '0.75rem 1rem',
                        border: 'none',
                        borderRadius: '8px',
                        fontSize: typography.sizes.base,
                        fontFamily: typography.fonts.primary,
                        backgroundColor: colors.background.primary,
                        color: colors.text.primary,
                        width: '100%'
                    }}
                    panelStyle={{
                        backgroundColor: colors.background.primary,
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        color: colors.text.primary
                    }}
                />
                {formik.touched.pickupDate && formik.errors.pickupDate && (
                    <small style={getErrorStyle()}>
                        {formik.errors.pickupDate}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="deliveryDate" style={getLabelStyle()}>
                    Delivery Date
                </label>
                <Calendar
                    id="deliveryDate"
                    name="deliveryDate"
                    value={formik.values.deliveryDate}
                    onChange={(e) => formik.setFieldValue('deliveryDate', e.value)}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.deliveryDate && formik.errors.deliveryDate ? 'p-invalid' : ''}`}
                    placeholder="Select delivery date"
                    showIcon
                    showTime
                    style={{
                        backgroundColor: colors.background.primary,
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        width: '100%'
                    }}
                    inputStyle={{
                        padding: '0.75rem 1rem',
                        border: 'none',
                        borderRadius: '8px',
                        fontSize: typography.sizes.base,
                        fontFamily: typography.fonts.primary,
                        backgroundColor: colors.background.primary,
                        color: colors.text.primary,
                        width: '100%'
                    }}
                    panelStyle={{
                        backgroundColor: colors.background.primary,
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        color: colors.text.primary
                    }}
                />
                {formik.touched.deliveryDate && formik.errors.deliveryDate && (
                    <small style={getErrorStyle()}>
                        {formik.errors.deliveryDate}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="status" style={getLabelStyle()}>
                    Status
                </label>
                <Dropdown
                    id="status"
                    name="status"
                    value={formik.values.status}
                    onChange={(e) => formik.setFieldValue('status', e.value)}
                    onBlur={formik.handleBlur}
                    options={statusOptions}
                    placeholder="Select status"
                    className={`w-full ${formik.touched.status && formik.errors.status ? 'p-invalid' : ''}`}
                    style={{
                        backgroundColor: colors.background.primary,
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        fontSize: typography.sizes.base,
                        fontFamily: typography.fonts.primary,
                        width: '100%',
                        minHeight: '3rem'
                    }}
                    panelStyle={{
                        backgroundColor: colors.background.primary,
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        fontFamily: typography.fonts.primary,
                        fontSize: typography.sizes.base
                    }}
                />
                {formik.touched.status && formik.errors.status && (
                    <small style={getErrorStyle()}>
                        {formik.errors.status}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="totalWeight" style={getLabelStyle()}>
                    Total Weight (lbs)
                </label>
                <InputNumber
                    id="totalWeight"
                    name="totalWeight"
                    value={formik.values.totalWeight}
                    onValueChange={(e) => formik.setFieldValue('totalWeight', e.value)}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.totalWeight && formik.errors.totalWeight ? 'p-invalid' : ''}`}
                    placeholder="Enter total weight"
                    min={0}
                    inputStyle={{
                        padding: '0.75rem 1rem',
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        fontSize: typography.sizes.base,
                        fontFamily: typography.fonts.primary,
                        backgroundColor: colors.background.primary,
                        color: colors.text.primary,
                        width: '100%'
                    }}
                />
                {formik.touched.totalWeight && formik.errors.totalWeight && (
                    <small style={getErrorStyle()}>
                        {formik.errors.totalWeight}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="customerName" style={getLabelStyle()}>
                    Customer Name
                </label>
                <InputText
                    id="customerName"
                    name="customerName"
                    value={formik.values.customerName}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.customerName && formik.errors.customerName ? 'p-invalid' : ''}`}
                    placeholder="Enter customer name"
                    style={getInputTextStyle()}
                />
                {formik.touched.customerName && formik.errors.customerName && (
                    <small style={getErrorStyle()}>
                        {formik.errors.customerName}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="customerTMSId" style={getLabelStyle()}>
                    Customer TMS ID
                </label>
                <InputText
                    id="customerTMSId"
                    name="customerTMSId"
                    value={formik.values.customerTMSId}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.customerTMSId && formik.errors.customerTMSId ? 'p-invalid' : ''}`}
                    placeholder="Enter customer TMS ID"
                    style={getInputTextStyle()}
                />
                {formik.touched.customerTMSId && formik.errors.customerTMSId && (
                    <small style={getErrorStyle()}>
                        {formik.errors.customerTMSId}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="minTemp" style={getLabelStyle()}>
                    Min Temperature (°F)
                </label>
                <InputNumber
                    id="minTemp"
                    name="minTemp"
                    value={formik.values.minTemp}
                    onValueChange={(e) => formik.setFieldValue('minTemp', e.value)}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.minTemp && formik.errors.minTemp ? 'p-invalid' : ''}`}
                    placeholder="Enter minimum temperature"
                    inputStyle={{
                        padding: '0.75rem 1rem',
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        fontSize: typography.sizes.base,
                        fontFamily: typography.fonts.primary,
                        backgroundColor: colors.background.primary,
                        color: colors.text.primary,
                        width: '100%'
                    }}
                />
                {formik.touched.minTemp && formik.errors.minTemp && (
                    <small style={getErrorStyle()}>
                        {formik.errors.minTemp}
                    </small>
                )}
            </div>
            <div>
                <label htmlFor="maxTemp" style={getLabelStyle()}>
                    Max Temperature (°F)
                </label>
                <InputNumber
                    id="maxTemp"
                    name="maxTemp"
                    value={formik.values.maxTemp}
                    onValueChange={(e) => formik.setFieldValue('maxTemp', e.value)}
                    onBlur={formik.handleBlur}
                    className={`w-full ${formik.touched.maxTemp && formik.errors.maxTemp ? 'p-invalid' : ''}`}
                    placeholder="Enter maximum temperature"
                    inputStyle={{
                        padding: '0.75rem 1rem',
                        border: `1px solid ${colors.border.light}`,
                        borderRadius: '8px',
                        fontSize: typography.sizes.base,
                        fontFamily: typography.fonts.primary,
                        backgroundColor: colors.background.primary,
                        color: colors.text.primary,
                        width: '100%'
                    }}
                />
                {formik.touched.maxTemp && formik.errors.maxTemp && (
                    <small style={getErrorStyle()}>
                        {formik.errors.maxTemp}
                    </small>
                )}
            </div>
        </div>
    );

    return (
        <Dialog
            header="Create New Load"
            visible={visible}
            onHide={handleDialogClose}
            style={{ 
                width: '50vw', 
                minWidth: '600px',
                background: colors.background.primary,
                borderRadius: '12px',
                boxShadow: '0 20px 60px rgba(0, 0, 0, 0.3)'
            }}
            contentStyle={{
                background: colors.background.primary,
                color: colors.text.primary,
                fontFamily: typography.fonts.primary,
                padding: '1.5rem'
            }}
            headerStyle={{
                background: colors.background.primary,
                color: colors.text.primary,
                borderBottom: `1px solid ${colors.border.light}`,
                borderRadius: '12px 12px 0 0',
                fontFamily: typography.fonts.primary,
                fontSize: typography.sizes.xl,
                fontWeight: typography.weights.semibold,
                padding: '1.5rem'
            }}
            modal
            draggable={false}
            resizable={false}
            position="center"
            maskStyle={{
                backgroundColor: colors.background.overlay
            }}
            footer={
                <div style={{ 
                    display: 'flex', 
                    justifyContent: 'space-between', 
                    gap: '1rem',
                    background: colors.background.primary,
                    padding: '1.5rem'
                }}>
                    <div>
                        {currentStep === 2 && (
                            <Button
                                label="Previous"
                                icon="pi pi-chevron-left"
                                onClick={handlePrevious}
                                style={{ 
                                    borderRadius: '10px',
                                    border: `2px solid ${colors.border.light}`,
                                    background: 'transparent',
                                    color: colors.text.secondary,
                                    fontFamily: typography.fonts.primary,
                                    fontSize: typography.sizes.base,
                                    fontWeight: typography.weights.semibold,
                                    padding: '0.75rem 1.5rem',
                                    transition: 'all 0.2s ease',
                                    boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)'
                                }}
                                onMouseEnter={(e) => {
                                    e.currentTarget.style.background = colors.background.tertiary;
                                    e.currentTarget.style.borderColor = colors.border.medium;
                                    e.currentTarget.style.transform = 'translateY(-1px)';
                                    e.currentTarget.style.boxShadow = '0 4px 8px rgba(0, 0, 0, 0.15)';
                                }}
                                onMouseLeave={(e) => {
                                    e.currentTarget.style.background = 'transparent';
                                    e.currentTarget.style.borderColor = colors.border.light;
                                    e.currentTarget.style.transform = 'translateY(0)';
                                    e.currentTarget.style.boxShadow = '0 2px 4px rgba(0, 0, 0, 0.1)';
                                }}
                            />
                        )}
                    </div>
                    <Button
                        label={currentStep === 1 ? 'Next' : 'Submit'}
                        icon={currentStep === 1 ? 'pi pi-chevron-right' : 'pi pi-check'}
                        iconPos={currentStep === 1 ? 'right' : 'left'}
                        onClick={handleNext}
                        loading={isLoading}
                        disabled={isLoading}
                        style={{
                            background: `linear-gradient(135deg, ${colors.brand.primary} 0%, ${colors.brand.secondary} 100%)`,
                            border: 'none',
                            borderRadius: '10px',
                            fontFamily: typography.fonts.primary,
                            fontSize: typography.sizes.base,
                            fontWeight: typography.weights.semibold,
                            padding: '0.75rem 1.5rem',
                            color: 'white',
                            transition: 'all 0.2s ease',
                            boxShadow: '0 4px 12px rgba(25, 118, 210, 0.3)',
                            opacity: isLoading ? 0.7 : 1
                        }}
                        onMouseEnter={(e) => {
                            if (!isLoading) {
                                e.currentTarget.style.transform = 'translateY(-2px)';
                                e.currentTarget.style.boxShadow = '0 6px 16px rgba(25, 118, 210, 0.4)';
                            }
                        }}
                        onMouseLeave={(e) => {
                            if (!isLoading) {
                                e.currentTarget.style.transform = 'translateY(0)';
                                e.currentTarget.style.boxShadow = '0 4px 12px rgba(25, 118, 210, 0.3)';
                            }
                        }}
                    />
                </div>
            }
        >
            <div style={{ 
                marginBottom: '2rem',
                background: colors.background.primary
            }}>
                <ProgressBar 
                    value={currentStep * 50} 
                    style={{ 
                        height: '4px',
                        borderRadius: '2px',
                        backgroundColor: colors.background.tertiary,
                        border: 'none'
                    }}
                    color={colors.brand.primary}
                />
            </div>

            <form onSubmit={formik.handleSubmit} style={{ background: colors.background.primary }}>
                {currentStep === 1 ? renderStep1() : renderStep2()}
            </form>
        </Dialog>
    );
};

export default CreateLoadDialog;
