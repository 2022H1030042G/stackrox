import React, { ReactElement } from 'react';
import { TextInput, PageSection, Form } from '@patternfly/react-core';

import * as yup from 'yup';

import { ClusterInitBundle } from 'services/ClustersService';
import usePageState from 'Containers/Integrations/hooks/usePageState';
import NotFoundMessage from 'Components/NotFoundMessage';
import useIntegrationForm from '../../useIntegrationForm';
import IntegrationFormActions from '../../IntegrationFormActions';
import FormCancelButton from '../../FormCancelButton';
import FormSaveButton from '../../FormSaveButton';
import ClusterInitBundleFormMessageAlert, {
    ClusterInitBundleFormResponseMessage,
} from './ClusterInitBundleFormMessageAlert';
import FormLabelGroup from '../../FormLabelGroup';
import ClusterInitBundleDetails from './ClusterInitBundleDetails';

export type ClusterInitBundleIntegration = ClusterInitBundle;

export type ClusterInitBundleIntegrationFormValues = {
    name: string;
};

export type ClusterInitBundleIntegrationFormProps = {
    initialValues: ClusterInitBundleIntegration | null;
    isEditable?: boolean;
};

export const validationSchema = yup.object().shape({
    name: yup.string().required('Required'),
});

export const defaultValues: ClusterInitBundleIntegrationFormValues = {
    name: '',
};

function ClusterInitBundleIntegrationForm({
    initialValues = null,
    isEditable = false,
}: ClusterInitBundleIntegrationFormProps): ReactElement {
    const formInitialValues = initialValues ? { ...initialValues, defaultValues } : defaultValues;
    const {
        values,
        errors,
        setFieldValue,
        isSubmitting,
        isTesting,
        onSave,
        onCancel,
        message,
    } = useIntegrationForm<ClusterInitBundleIntegrationFormValues, typeof validationSchema>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const { isEditing, isViewingDetails } = usePageState();
    const isGenerated = Boolean(message?.responseData);

    function onChange(value, event) {
        return setFieldValue(event.target.id, value, false);
    }

    // The edit flow doesn't make sense for Cluster Init Bundles so we'll show an empty state message here
    if (isEditing) {
        return (
            <NotFoundMessage
                title="This Cluster Init Bundle can not be edited"
                message="Create a new Cluster Init Bundle or delete an existing one"
            />
        );
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                {message && (
                    <div className="pf-u-pb-md">
                        <ClusterInitBundleFormMessageAlert
                            message={message as ClusterInitBundleFormResponseMessage}
                        />
                    </div>
                )}
                {isViewingDetails && initialValues ? (
                    <ClusterInitBundleDetails meta={initialValues} />
                ) : (
                    <Form isWidthLimited>
                        <FormLabelGroup
                            label="Cluster Init Bundle Name"
                            isRequired
                            fieldId="name"
                            errors={errors}
                        >
                            <TextInput
                                isRequired
                                type="text"
                                id="name"
                                name="name"
                                value={values.name}
                                onChange={onChange}
                                isDisabled={!isEditable || isGenerated}
                            />
                        </FormLabelGroup>
                    </Form>
                )}
            </PageSection>
            {isEditable &&
                (!isGenerated ? (
                    <IntegrationFormActions>
                        <FormSaveButton
                            onSave={onSave}
                            isSubmitting={isSubmitting}
                            isTesting={isTesting}
                        >
                            Generate
                        </FormSaveButton>
                        <FormCancelButton onCancel={onCancel} isDisabled={isSubmitting}>
                            Cancel
                        </FormCancelButton>
                    </IntegrationFormActions>
                ) : (
                    <IntegrationFormActions>
                        <FormCancelButton onCancel={onCancel} isDisabled={isSubmitting}>
                            Close
                        </FormCancelButton>
                    </IntegrationFormActions>
                ))}
        </>
    );
}

export default ClusterInitBundleIntegrationForm;