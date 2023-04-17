///
//  Generated code. Do not modify.
//  source: api/auth/auth.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use secretDescriptor instead')
const Secret$json = const {
  '1': 'Secret',
  '2': const [
    const {'1': 'login', '3': 1, '4': 1, '5': 9, '10': 'login'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
    const {'1': 'role', '3': 3, '4': 1, '5': 9, '10': 'role'},
  ],
};

/// Descriptor for `Secret`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List secretDescriptor = $convert.base64Decode('CgZTZWNyZXQSFAoFbG9naW4YASABKAlSBWxvZ2luEhoKCHBhc3N3b3JkGAIgASgJUghwYXNzd29yZBISCgRyb2xlGAMgASgJUgRyb2xl');
@$core.Deprecated('Use loginRequestDescriptor instead')
const LoginRequest$json = const {
  '1': 'LoginRequest',
  '2': const [
    const {'1': 'secret', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_auth_service.api.auth.Secret', '10': 'secret'},
  ],
};

/// Descriptor for `LoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginRequestDescriptor = $convert.base64Decode('CgxMb2dpblJlcXVlc3QSUAoGc2VjcmV0GAEgASgLMjguZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfYXV0aF9zZXJ2aWNlLmFwaS5hdXRoLlNlY3JldFIGc2VjcmV0');
@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse$json = const {
  '1': 'LoginResponse',
  '2': const [
    const {'1': 'access_token', '3': 1, '4': 1, '5': 9, '10': 'accessToken'},
  ],
};

/// Descriptor for `LoginResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginResponseDescriptor = $convert.base64Decode('Cg1Mb2dpblJlc3BvbnNlEiEKDGFjY2Vzc190b2tlbhgBIAEoCVILYWNjZXNzVG9rZW4=');
@$core.Deprecated('Use signUpRequestDescriptor instead')
const SignUpRequest$json = const {
  '1': 'SignUpRequest',
  '2': const [
    const {'1': 'secret', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_auth_service.api.auth.Secret', '10': 'secret'},
  ],
};

/// Descriptor for `SignUpRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signUpRequestDescriptor = $convert.base64Decode('Cg1TaWduVXBSZXF1ZXN0ElAKBnNlY3JldBgBIAEoCzI4LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2F1dGhfc2VydmljZS5hcGkuYXV0aC5TZWNyZXRSBnNlY3JldA==');
@$core.Deprecated('Use signUpResponseDescriptor instead')
const SignUpResponse$json = const {
  '1': 'SignUpResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `SignUpResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List signUpResponseDescriptor = $convert.base64Decode('Cg5TaWduVXBSZXNwb25zZRIOCgJpZBgBIAEoA1ICaWQ=');
const $core.Map<$core.String, $core.dynamic> AuthServiceBase$json = const {
  '1': 'Auth',
  '2': const [
    const {'1': 'Login', '2': '.github.constantine27k.crnt_auth_service.api.auth.LoginRequest', '3': '.github.constantine27k.crnt_auth_service.api.auth.LoginResponse', '4': const {}},
    const {'1': 'SignUp', '2': '.github.constantine27k.crnt_auth_service.api.auth.SignUpRequest', '3': '.github.constantine27k.crnt_auth_service.api.auth.SignUpResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use authServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AuthServiceBase$messageJson = const {
  '.github.constantine27k.crnt_auth_service.api.auth.LoginRequest': LoginRequest$json,
  '.github.constantine27k.crnt_auth_service.api.auth.Secret': Secret$json,
  '.github.constantine27k.crnt_auth_service.api.auth.LoginResponse': LoginResponse$json,
  '.github.constantine27k.crnt_auth_service.api.auth.SignUpRequest': SignUpRequest$json,
  '.github.constantine27k.crnt_auth_service.api.auth.SignUpResponse': SignUpResponse$json,
};

/// Descriptor for `Auth`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List authServiceDescriptor = $convert.base64Decode('CgRBdXRoEp4BCgVMb2dpbhI+LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2F1dGhfc2VydmljZS5hcGkuYXV0aC5Mb2dpblJlcXVlc3QaPy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9hdXRoX3NlcnZpY2UuYXBpLmF1dGguTG9naW5SZXNwb25zZSIUgtPkkwIOOgEqIgkvdjEvbG9naW4SowEKBlNpZ25VcBI/LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2F1dGhfc2VydmljZS5hcGkuYXV0aC5TaWduVXBSZXF1ZXN0GkAuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfYXV0aF9zZXJ2aWNlLmFwaS5hdXRoLlNpZ25VcFJlc3BvbnNlIhaC0+STAhA6ASoiCy92MS9zaWduX3Vw');
